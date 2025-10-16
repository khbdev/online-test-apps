package handler

import (
	testpb "github.com/khbdev/proto-online-test/proto/test"
	"test-section-service/internal/service"
)

type TestHandler struct {
	testpb.UnimplementedTestServiceServer
	SectionService  *service.SectionService
	QuestionService *service.QuestionService
	OptionService   *service.OptionService
}

func NewTestHandler(sectionService *service.SectionService, questionService *service.QuestionService, optionService *service.OptionService) *TestHandler {
	return &TestHandler{
		SectionService:  sectionService,
		QuestionService: questionService,
		OptionService:   optionService,
	}
}
