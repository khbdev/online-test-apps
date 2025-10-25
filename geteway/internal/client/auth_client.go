package client

import (
	"context"
	"fmt"
	"geteway-service/internal/discovery"
	"log"
	"time"

	 authpb "github.com/khbdev/proto-online-test/proto/auth"
	"google.golang.org/grpc"
)

type AuthClient struct {
	conn   *grpc.ClientConn
	client authpb.AuthServiceClient
}


func NewAuthClient() (*AuthClient, error) {
	addr, err := discovery.GetServiceAddress("auth-service")
	if err != nil {
		return nil, fmt.Errorf("auth service address topilmadi: %v", err)
	}

	var conn *grpc.ClientConn
	var dialErr error

	for i := 1; i <= 3; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		conn, dialErr = grpc.DialContext(ctx, addr, grpc.WithInsecure(), grpc.WithBlock())
		if dialErr == nil {
			log.Printf("[AuthClient] ✅ Ulandi: %s", addr)
			client := authpb.NewAuthServiceClient(conn)
			return &AuthClient{conn: conn, client: client}, nil
		}

		log.Printf("[AuthClient] ❌ Ulanib bo‘lmadi (urinish %d/3): %v", i, dialErr)
		time.Sleep(1 * time.Second)
	}

	return nil, fmt.Errorf("auth service bilan 3 marta ulanish muvaffaqiyatsiz: %v", dialErr)
}


func (a *AuthClient) Close() {
	if a.conn != nil {
		_ = a.conn.Close()
	}
}


func (a *AuthClient) Login(username, password string) (*authpb.LoginResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req := &authpb.LoginRequest{
		Username: username,
		Password: password,
	}

	res, err := a.client.Login(ctx, req)
	if err != nil {
		log.Printf("[AuthClient] Login xatolik: %v", err)
		return nil, err
	}

	log.Printf("[AuthClient]  Login muvaffaqiyatli: %s", username)
	return res, nil
}
