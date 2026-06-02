package main

import (
	"net/http"
	"tickets/internal/user"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CustomValidator struct {
  validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
  return cv.validator.Struct(i)
}

func main() {
dsn := "host=localhost user=postgres password=rahat12 dbname=tickets port=9920 sslmode=disable TimeZone=Asia/Shanghai"
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})	
    if err != nil {
	   panic("failed to connect database")
    }

    if err := db.AutoMigrate(&user.User{}); err != nil {
	   panic("failed to migrate database")
   }
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(middleware.RequestLogger())

	e.GET("/", func(c *echo.Context) error {
		return c.String(http.StatusOK, "Tickets counter!")
	})

	e.Validator = &CustomValidator{validator: validator.New()}

	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	e.POST("/users", userHandler.CreateUser)

	if err := e.Start(":8080"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}

}
