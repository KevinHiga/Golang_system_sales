package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		dynamicCORSConfig := middleware.CORSConfig{
			AllowOrigins:     []string{"*", "*"},
			AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
			AllowCredentials: true,
		}
		CORSMiddleware := middleware.CORSWithConfig(dynamicCORSConfig)
		CORSHandler := CORSMiddleware(next)
		return CORSHandler(c)
	}
}
