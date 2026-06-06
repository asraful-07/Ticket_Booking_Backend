package user

import (
	"net/http"
	htttprespones "tickets/internal/htttpRespones"
	"tickets/internal/user/dto"

	"github.com/labstack/echo/v5"
)

type handler struct {
	service *service
}

func NewHandler(service *service) *handler {
	return &handler{service: service}
}

func (h *handler) CreateUser(c *echo.Context) error {
	var req dto.CreateUserRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, htttprespones.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid request payload",
			Details: err.Error(),
		})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, htttprespones.Error{
			Code:    http.StatusBadRequest,
			Message: "Validation failed",
			Details: err.Error(),
		})
	}

    response, err := h.service.CreateUser(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, htttprespones.Error{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create user",
			Details: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, response)
}


func (h *handler) LoginUser(c *echo.Context) error {
	var req dto.LoginRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, htttprespones.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid request payload",
			Details: err.Error(),
		})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, htttprespones.Error{
			Code:    http.StatusBadRequest,
			Message: "Validation failed",
			Details: err.Error(),
		})
	}

	response, err := h.service.LoginUser(req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, htttprespones.Error{
			Code:    http.StatusUnauthorized,
			Message: "Invalid credentials",
			Details: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response)
}

func (h *handler) GetMe(c *echo.Context) error {
	userID, ok := c.Get("user_id").(uint)
	if !ok {
		return c.JSON(http.StatusUnauthorized, htttprespones.Error{
			Code:    http.StatusUnauthorized,
			Message: "Cannot get user information",
			Details: "missing user id in context",
		})
	}

	email, _ := c.Get("user_email").(string)
	name, _ := c.Get("user_name").(string)

	return c.JSON(http.StatusOK, dto.Response{
		Id:    userID,
		Name:  name,
		Email: email,
	})
}
