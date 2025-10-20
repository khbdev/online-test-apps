package main

import (
	"log"
	"net"

	"test-generation-servis/internal/config"
	"test-generation-servis/internal/client"
	"test-generation-servis/internal/handler"
	repo "test-generation-servis/internal/repostory"
	"test-generation-servis/internal/service"

	pb "github.com/khbdev/proto-online-test/proto/generate"
	"google.golang.org/grpc"
)

func main() {

	config.InitRedis()


	testRepo := repo.NewRepository(config.RedisClient)


	sectionClient := client.NewSectionClient("127.0.0.1:50054") 

	testService := service.NewTestService(testRepo, sectionClient, "http://localhost:8080") 


	testHandler := handler.NewTestHandler(testService)


	lis, err := net.Listen("tcp", ":50055") 
	if err != nil {
		log.Fatalf("Port band: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTestServiceServer(grpcServer, testHandler)

	log.Println("Test Generation Service running on port :50055")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Serverda xatolik: %v", err)
	}
}
