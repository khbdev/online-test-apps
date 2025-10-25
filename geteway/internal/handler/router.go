package handler

import (
	"geteway-service/internal/client"
	"geteway-service/internal/service"
	"github.com/gin-gonic/gin"
	"log"
)

func SetupRoutes(r *gin.Engine) {
	
	authClient, err := client.NewAuthClient()
	if err != nil {
		log.Fatalf("[Routes] AuthClient ulanishda xatolik: %v", err)
	}


	authService := service.NewAuthService(authClient)
	authHandler := NewAuthHandler(authService)


	auth := r.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)

	}
}
