package handler

import (
	"context"
	"test-section-service/internal/models"
	"test-section-service/internal/service"

	pb "github.com/khbdev/proto-online-test/proto/test"
	"google.golang.org/protobuf/types/known/emptypb"
)

type SectionHandler struct {
	pb.UnimplementedTestServiceServer
	sectionService *service.SectionService
}

func NewSectionHandler(sectionService *service.SectionService) *SectionHandler {
	return &SectionHandler{sectionService: sectionService}
}


func (h *SectionHandler) CreateSection(ctx context.Context, req *pb.CreateSectionRequest) (*pb.Section, error) {
	section := models.Section{Name: req.Name}
	if err := h.sectionService.CreateSection(&section); err != nil {
		return nil, err
	}

	return &pb.Section{
		Id:   uint64(section.ID),
		Name: section.Name,
	}, nil
}


func (h *SectionHandler) GetSectionById(ctx context.Context, req *pb.SectionID) (*pb.GetSectionResponse, error) {
	section, err := h.sectionService.GetSectionByID(uint(req.Id))
	if err != nil {
		return nil, err
	}

	return &pb.GetSectionResponse{
		Section: &pb.Section{
			Id:   uint64(section.ID),
			Name: section.Name,
		},
	}, nil
}


func (h *SectionHandler) GetAllSections(ctx context.Context, _ *emptypb.Empty) (*pb.GetAllSectionsResponse, error) {
	sections, err := h.sectionService.GetAllSections()
	if err != nil {
		return nil, err
	}

	var pbSections []*pb.Section
	for _, s := range sections {
		pbSections = append(pbSections, &pb.Section{
			Id:   uint64(s.ID),
			Name: s.Name,
		})
	}

	return &pb.GetAllSectionsResponse{Sections: pbSections}, nil
}


func (h *SectionHandler) UpdateSection(ctx context.Context, req *pb.UpdateSectionRequest) (*pb.Section, error) {
	section, err := h.sectionService.GetSectionByID(uint(req.Id))
	if err != nil {
		return nil, err
	}

	section.Name = req.Name
	if err := h.sectionService.UpdateSection(section); err != nil {
		return nil, err
	}

	return &pb.Section{
		Id:   uint64(section.ID),
		Name: section.Name,
	}, nil
}


func (h *SectionHandler) DeleteSection(ctx context.Context, req *pb.SectionID) (*emptypb.Empty, error) {
	if err := h.sectionService.DeleteSection(uint(req.Id)); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}


func (h *SectionHandler) GetFullSectionStructure(ctx context.Context, req *pb.GetFullSectionRequest) (*pb.GetFullSectionResponse, error) {
	section, err := h.sectionService.GetFullSectionStructure(uint(req.SectionId))
	if err != nil {
		return nil, err
	}

	var pbQuestions []*pb.Question
	for _, q := range section.Questions {
		var pbOptions []*pb.Option
		for _, o := range q.Options {
			pbOptions = append(pbOptions, &pb.Option{
				Id:         uint64(o.ID),
				Text:       o.Text,
				IsCorrect:  o.IsCorrect,
				QuestionId: uint64(o.QuestionID),
			})
		}

		pbQuestions = append(pbQuestions, &pb.Question{
			Id:        uint64(q.ID),
			Text:      q.Text,
			SectionId: uint64(q.SectionID),
			Options:   pbOptions,
		})
	}

	return &pb.GetFullSectionResponse{
		Section: &pb.Section{
			Id:        uint64(section.ID),
			Name:      section.Name,
			Questions: pbQuestions,
		},
	}, nil
}
