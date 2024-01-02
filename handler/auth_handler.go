package handler

import (
	"net/http"

	"github.com/Siravitt/go-todoist/logs"
	"github.com/Siravitt/go-todoist/service/auth_service"
	"github.com/labstack/echo/v4"
)

type authHandler struct {
	authSrv auth_service.AuthService
}

func NewAuthHandler(authSrv auth_service.AuthService) authHandler {
	return authHandler{authSrv: authSrv}
}

func (h authHandler) Login(c echo.Context) error {
	loginRequest := auth_service.LoginRequest{}
	err := c.Bind(&loginRequest)
	if err != nil {
		logs.Error(err)
		return handleError(c, err)
	}
	token, err := h.authSrv.Login(loginRequest)
	if err != nil {
		logs.Error(err)
		return handleError(c, err)
	}
	return c.JSON(http.StatusOK, token)
}

func (h authHandler) Register(c echo.Context) error {
	registerRequest := auth_service.RegisterRequest{}
	err := c.Bind(&registerRequest)
	if err != nil {
		logs.Error(err)
		return handleError(c, err)
	}
	result, err := h.authSrv.Register(registerRequest)
	if err != nil {
		logs.Error(err)
		return handleError(c, err)
	}
	return c.JSON(http.StatusOK, result)
}
