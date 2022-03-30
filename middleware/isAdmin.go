package middleware

import (
	"log"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		admin := claims["admin"].(bool)
		log.Printf("isAdmin %v", admin)
		if admin {
			next(c)
		}
		return echo.ErrUnauthorized
	}
}
