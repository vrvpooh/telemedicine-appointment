package route

import (
	"telemedicine-api/handler"
	"telemedicine-api/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine) {
	api := r.Group("/api")

	// Public routes
	auth := api.Group("/auth")
	{
		auth.POST("/register", handler.RegisterUser)
		auth.POST("/login", handler.LoginUser)
	}

	// Protected routes (ต้องใส่ Token)
	users := api.Group("/users")
	users.Use(middleware.AuthMiddleware()) // เรียกใช้ Middleware กั้นตรงนี้
	{
		users.GET("/me", handler.GetMe)
	}
}
