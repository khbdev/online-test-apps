package handler

import (
	"context"

	"test-section-service/internal/models"
	"test-section-service/internal/service"

	pb "github.com/khbdev/proto-online-test/proto/test"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

type QuestionHandler struct {
	pb.UnimplementedTestServiceServer
	questionService *service.QuestionService
}

func NewQuestionHandler(s *service.QuestionService) *QuestionHandler {
	return &QuestionHandler{questionService: s}
}


func (h *QuestionHandler) CreateQuestion(ctx context.Context, req *pb.CreateQuestionRequest) (*pb.Question, error) {
	question := &models.Question{
		SectionID: uint(req.SectionId),
		Text:      req.Text,
	}

	created, err := h.questionService.Create(question)
	if err != nil {
		return nil, err
	}

	return &pb.Question{
		Id:        uint64(created.ID),
		SectionId: uint64(created.SectionID),
		Text:      created.Text,
	}, nil
}


func (h *QuestionHandler) GetQuestionById(ctx context.Context, req *pb.QuestionID) (*pb.GetQuestionResponse, error) {
	q, err := h.questionService.GetByID(uint(req.Id))
	if err != nil {
		return nil, err
	}

	return &pb.GetQuestionResponse{
		Question: &pb.Question{
			Id:        uint64(q.ID),
			SectionId: uint64(q.SectionID),
			Text:      q.Text,
		},
	}, nil
}


func (h *QuestionHandler) GetAllQuestions(ctx context.Context, _ *emptypb.Empty) (*pb.GetAllQuestionsResponse, error) {
	questions, err := h.questionService.GetAll()
	if err != nil {
		return nil, err
	}

	var list []*pb.Question
	for _, q := range questions {
		list = append(list, &pb.Question{
			Id:        uint64(q.ID),
			SectionId: uint64(q.SectionID),
			Text:      q.Text,
		})
	}

	return &pb.GetAllQuestionsResponse{Questions: list}, nil
}


func (h *QuestionHandler) UpdateQuestion(ctx context.Context, req *pb.UpdateQuestionRequest) (*pb.Question, error) {
	q := &models.Question{
		Model:     gorm.Model{ID: uint(req.Id)},
		SectionID: uint(req.SectionId),
		Text:      req.Text,
	}

	updated, err := h.questionService.Update(q)
	if err != nil {
		return nil, err
	}

	return &pb.Question{
		Id:        uint64(updated.ID),
		SectionId: uint64(updated.SectionID),
		Text:      updated.Text,
	}, nil
}


func (h *QuestionHandler) DeleteQuestion(ctx context.Context, req *pb.QuestionID) (*emptypb.Empty, error) {
	if err := h.questionService.Delete(uint(req.Id)); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
