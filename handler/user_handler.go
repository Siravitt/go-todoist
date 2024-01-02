package handler

import (
	"net/http"
	"strconv"

	"github.com/Siravitt/go-todoist/logs"
	"github.com/Siravitt/go-todoist/service/user_service"
	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userSrv user_service.UserService
}

func NewUserHandler(userSrv user_service.UserService) userHandler {
	return userHandler{userSrv: userSrv}
}

func (h userHandler) GetUsers(c echo.Context) error {
	users, err := h.userSrv.GetUsers()
	if err != nil {
		logs.Error(err)
		return handleError(c, err)
	}
	return c.JSON(http.StatusOK, users)
}

func (h userHandler) GetUser(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logs.Error(err)
		return handleError(c, err)
	}
	user, err := h.userSrv.GetUser(userID)
	if err != nil {
		logs.Error(err)
		return handleError(c, err)
	}
	return c.JSON(http.StatusOK, user)
}
