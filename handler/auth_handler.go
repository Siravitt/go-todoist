package handler

import (
	"net/http"

	"github.com/Siravitt/go-todoist/service/auth_service"
	"github.com/Siravitt/go-todoist/utils"
	"github.com/labstack/echo/v4"
)

type authHandler struct {
	authSrv auth_service.AuthService
}

func NewAuthHandler(authSrv auth_service.AuthService) authHandler {
	return authHandler{authSrv: authSrv}
}

func (h authHandler) Login(c echo.Context) error {
	token, _ := utils.GenerateJWT(1)
	return c.JSON(http.StatusOK, token)
}

func (h authHandler) Register(c echo.Context) error {
	return nil
}
