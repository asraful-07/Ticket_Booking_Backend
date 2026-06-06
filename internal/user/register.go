package user

import (
	"tickets/internal/auth"
	"tickets/internal/middlewares"
	"time"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB) {
	userRepo := NewRepository(db)
	jwtService := auth.NewJWTService("BETTER_AUTH_SECRET=Sy9frXLrOngAQXcMiuF7yAfmNTUgziBH", 24*time.Hour)
	userService := NewService(userRepo, jwtService)
	userHandler := NewHandler(userService)

    api := e.Group("/api/v1/auth")

	api.POST("/register", userHandler.CreateUser)
	api.POST("/login", userHandler.LoginUser)
	api.GET("/me", userHandler.GetMe, middlewares.AuthMiddleware(jwtService))
}