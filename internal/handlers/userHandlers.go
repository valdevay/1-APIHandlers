package handlers

import (
	"internal/userService"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service userService.UserService
}

func (h *UserHandler) GetUsers(c echo.Context) error {
	// реализация
}

func (h *UserHandler) PostUser(c echo.Context) error {
	// реализация
}

func (h *UserHandler) PatchUserByID(c echo.Context) error {
	// реализация
}

func (h *UserHandler) DeleteUserByID(c echo.Context) error {
	// реализация
}
