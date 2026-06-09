package server

import (
	"tickets/internal/booking"
	"tickets/internal/config"
	"tickets/internal/event"
	"tickets/internal/user"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"gorm.io/gorm"
)

type CustomValidator struct {
  validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
  return cv.validator.Struct(i)
}

func StartServer( db *gorm.DB, cfg *config.Config) {
	if err := db.AutoMigrate(&user.User{},  &event.Event{}); err != nil {
		panic("failed to migrate database")
	}

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	e.Use(middleware.RequestLogger())

	// All Routes
	user.RegisterRoutes(e, db, cfg)
	event.RegisterRoutes(e, db)
    booking.RegisterRoutes(e, db, cfg)

	if err := e.Start(":" + cfg.PORT); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}