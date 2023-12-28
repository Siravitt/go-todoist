package auth_middleware

import (
	"net/http"
	"strings"

	"github.com/Siravitt/go-todoist/utils"
	"github.com/labstack/echo/v4"
)

type AuthMiddleware interface {
	AuthorizationMiddleware(echo.HandlerFunc) echo.HandlerFunc
}

type authMiddleware struct {
	c echo.Context
}

type unauthorizedResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func NewAuthMiddleware() AuthMiddleware {
	return authMiddleware{}
}

func (m authMiddleware) AuthorizationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		s := c.Request().Header.Get("Authorization")
		token := strings.TrimPrefix(s, "Bearer ")
		claim, err := utils.ValidateJWT(token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, unauthorizedResponse{Success: false, Message: "You're unauthorize"})
		}
		c.Set("userID", claim["UserID"])
		return next(c)
	}
}
