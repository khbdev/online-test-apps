package handler

import (
	"auth-service/internal/service"
	"context"

	 pb "github.com/khbdev/proto-online-test/proto/auth"
)

type AuthHandler struct {
	pb.UnimplementedAuthServiceServer
	svc *service.AuthService
}

func NewAuthHandler(s *service.AuthService) *AuthHandler {
	return &AuthHandler{svc: s}
}

func (h *AuthHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	tokenPair, err := h.svc.Login(req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    int64(900),
	}, nil
}
