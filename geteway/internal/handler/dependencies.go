package handler

import (
	"geteway-service/internal/client"
	"geteway-service/internal/service"
	"log"
)

type Dependencies struct {
	AuthHandler  *AuthHandler
	AdminHandler *AdminHandler
	UserHandler *UserHandler
	TestHandler *TestHandler
	JobHandler *JobHandler
	FilterHandler *FilterHandler
	
}

func InitDependencies() *Dependencies {
	//  Auth Client va Service
	authClient, err := client.NewAuthClient()
	if err != nil {
		log.Fatalf("[Dependencies] AuthClient ulanishda xatolik: %v", err)
	}

	authService := service.NewAuthService(authClient)
	authHandler := NewAuthHandler(authService)

	//  Admin Service
	adminService, err := service.NewAdminService()
	if err != nil {
		log.Fatalf("[Dependencies] AdminService ulanishda xatolik: %v", err)
	}
	adminHandler := NewAdminHandler(adminService)

	// User Service

	  userService, err := service.NewUserService()
	  if err != nil {
		log.Fatalf("[Dependencies] UserService ulanishda xatolik: %v", err)
	}
	userHandler := NewUserHandler(userService)

	// Test Generate Service

	generateClint, err := client.NewGenerateClient()
if err != nil {
		log.Fatalf("[Dependencies] AuthClient ulanishda xatolik: %v", err)
	}

	 generateService := service.NewGenerateService(generateClint)

	 generateHandler := NewGenerateHandler(generateService)
// Job Service 
jobClient, err := client.NewJobClient()
if err != nil {
		log.Fatalf("[Dependencies] JobClient ulanishda xatolik: %v", err)
	}
jobService := service.NewJobService(jobClient)
JobHandler := NewJobHandler(jobService)

// filter service 
filterClient, err  := client.NewFilterClient()
if err != nil {
		log.Fatalf("[Dependencies] JobClient ulanishda xatolik: %v", err)
	}
	filterService := service.NewFilterService(filterClient)
	filterHandler := NewFilterHandler(filterService)

	return &Dependencies{
		AuthHandler:  authHandler,
		AdminHandler: adminHandler,
		UserHandler: userHandler,
		TestHandler: generateHandler,
		JobHandler: JobHandler,
		FilterHandler: filterHandler,
		
	}
}
