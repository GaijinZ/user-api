package authentication

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		isAdmin := claims["admin"].(bool)

		if !isAdmin {
			return echo.ErrUnauthorized
		}

		return next(c)
	}
}
