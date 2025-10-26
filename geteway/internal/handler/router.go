package handler

import (
	"geteway-service/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	deps := InitDependencies()

	// 🔹 Auth routes (JWTsiz)
	auth := r.Group("/auth")
	{
		auth.POST("/login", deps.AuthHandler.Login)
	}

	admin := r.Group("/admin")
	admin.Use(middleware.JWTMiddleware())
	{
		admin.GET("/list", deps.AdminHandler.GetAll)
		admin.GET("/:id", deps.AdminHandler.GetByID)
		admin.POST("/create", deps.AdminHandler.Create)
		admin.PUT("/update/:id", deps.AdminHandler.Update)
		admin.DELETE("/delete/:id", deps.AdminHandler.Delete)
	}

	user := r.Group("/user")
	user.Use(middleware.JWTMiddleware())
	{
		user.GET("/list", deps.UserHandler.GetAll)
		user.GET("/:id", deps.UserHandler.GetByID)
		user.POST("/create", deps.UserHandler.Create)
		user.PUT("/update/:id", deps.UserHandler.Update)
		user.DELETE("/delete/:id", deps.UserHandler.Delete)
	}

	test := r.Group("/test")
	test.Use(middleware.JWTMiddleware())
{
	
	test.POST("/generate", deps.TestHandler.GenerateTest)

	
	test.GET("/get", deps.TestHandler.GetTest)
}
}
