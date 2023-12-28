package handler

import (
	"net/http"

	"github.com/Siravitt/go-todoist/errors"

	"github.com/labstack/echo/v4"
)

type errorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func handleError(c echo.Context, err error) error {
	switch e := err.(type) {
	case errors.AppError:
		return c.JSON(e.Code, errorResponse{Success: false, Message: e.Message})
	case error:
		return c.JSON(http.StatusInternalServerError, errorResponse{Success: false, Message: "Unexpected error"})
	}
	return nil
}
