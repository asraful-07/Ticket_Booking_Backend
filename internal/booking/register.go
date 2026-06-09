package booking

import (
	"time"

	"tickets/internal/auth"
	"tickets/internal/config"
	"tickets/internal/event"
	"tickets/internal/middlewares"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB, cfg *config.Config) {
	bookingRepo := NewRepository(db)
	eventRepo := event.NewRepository(db)

	svc := NewService(bookingRepo, eventRepo)
	handler := NewHandler(svc)

	jwtService := auth.NewJWTService(cfg.JWTSecretKey, 24*time.Hour)

	api := e.Group("/api/v1/bookings", middlewares.AuthMiddleware(jwtService))

	api.POST("", handler.CreateBooking)
	api.GET("/me", handler.GetMyBookings)

}
