package handler

import (
	"context"
	"test-section-service/internal/models"
	"test-section-service/internal/service"

	testpb "github.com/khbdev/proto-online-test/proto/test"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

type OptionHandler struct {
	testpb.UnimplementedTestServiceServer
	Service *service.OptionService
}

func NewOptionHandler(s *service.OptionService) *OptionHandler {
	return &OptionHandler{Service: s}
}

func (h *OptionHandler) CreateOption(ctx context.Context, req *testpb.CreateOptionRequest) (*testpb.Option, error) {
	option := &models.Option{
		Text:       req.Text,
		IsCorrect:  req.IsCorrect,
		QuestionID: uint(req.QuestionId),
	}

	if err := h.Service.CreateOption(option); err != nil {
		return nil, err
	}

	return &testpb.Option{
		Id:         uint64(option.ID),
		Text:       option.Text,
		IsCorrect:  option.IsCorrect,
		QuestionId: uint64(option.QuestionID),
	}, nil
}

func (h *OptionHandler) GetAllOptions(ctx context.Context, _ *emptypb.Empty) (*testpb.GetAllOptionsResponse, error) {
	options, err := h.Service.GetAllOptions()
	if err != nil {
		return nil, err
	}

	var res []*testpb.Option
	for _, o := range options {
		res = append(res, &testpb.Option{
			Id:         uint64(o.ID),
			Text:       o.Text,
			IsCorrect:  o.IsCorrect,
			QuestionId: uint64(o.QuestionID),
		})
	}

	return &testpb.GetAllOptionsResponse{Options: res}, nil
}

func (h *OptionHandler) GetOptionById(ctx context.Context, req *testpb.OptionID) (*testpb.GetOptionResponse, error) {
	option, err := h.Service.GetOptionByID(uint(req.Id))
	if err != nil {
		return nil, err
	}

	return &testpb.GetOptionResponse{
		Option: &testpb.Option{
			Id:         uint64(option.ID),
			Text:       option.Text,
			IsCorrect:  option.IsCorrect,
			QuestionId: uint64(option.QuestionID),
		},
	}, nil
}

func (h *OptionHandler) UpdateOption(ctx context.Context, req *testpb.UpdateOptionRequest) (*testpb.Option, error) {
		option := &models.Option{
		Model: gorm.Model{ID: uint(req.Id)},
		Text:       req.Text,
		IsCorrect:  req.IsCorrect,
		QuestionID: uint(req.QuestionId),
	}

	if err := h.Service.UpdateOption(option); err != nil {
		return nil, err
	}

	return &testpb.Option{
		Id:         uint64(option.ID),
		Text:       option.Text,
		IsCorrect:  option.IsCorrect,
		QuestionId: uint64(option.QuestionID),
	}, nil
}

func (h *OptionHandler) DeleteOption(ctx context.Context, req *testpb.OptionID) (*emptypb.Empty, error) {
	if err := h.Service.DeleteOption(uint(req.Id)); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
