package service

import (
	"context"

	"fmt"

	adapter "geteway-service/internal/adabter"
	"geteway-service/internal/client"

	authpb "github.com/khbdev/proto-online-test/proto/auth"
)

type AuthService struct {
	authClient *client.AuthClient
}

func NewAuthService(authClient *client.AuthClient) *AuthService {
	return &AuthService{authClient: authClient}
}

func (s *AuthService) Login(ctx context.Context, body []byte) (map[string]interface{}, error) {
	// 🔹 REST → Proto
	protoReq, err := adapter.ProtoGenerate(body, &authpb.LoginRequest{})
	if err != nil {
		return nil, fmt.Errorf("Login: ProtoGenerate xatolik: %v", err)
	}

	req, ok := protoReq.(*authpb.LoginRequest)
	if !ok {
		return nil, fmt.Errorf("Login: noto‘g‘ri request turi")
	}

	// 🔹 AuthService (gRPC) orqali login
	res, err := s.authClient.Login(req.Username, req.Password)
	if err != nil {
		return nil, fmt.Errorf("Login: AuthClient orqali login xatolik: %v", err)
	}

	// 🔹 Proto → REST
	restRes, err := adapter.RestGenerate(res)
	if err != nil {
		return nil, fmt.Errorf("Login: RestGenerate xatolik: %v", err)
	}

	return restRes, nil
}
