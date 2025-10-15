package main

import (
	"fmt"
	"log"
	"net"
	"user-service/internal/config"
	"user-service/internal/handler"
	"user-service/internal/repostory"
	"user-service/internal/service"

	userpb "github.com/khbdev/proto-online-test/proto/user"
	"google.golang.org/grpc"
)

func main(){
	config.InitDB()
 
	repo := repostory.NewUserRepository(config.DB)
	service := service.NewUserService(repo)
	 handler := handler.NewUserHandler(service)


	   lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatal("Port band:", err)
	}
	s := grpc.NewServer()
	userpb.RegisterUserServiceServer(s, handler)

	fmt.Println("Admin gRPC server is running on port 50053")

	if err := s.Serve(lis); err != nil {
		log.Fatal("Server Error: ", err)
	}

}