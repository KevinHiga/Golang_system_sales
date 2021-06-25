package middleware

import (
	"fmt"
	"golang-project/config/security"
	sessionRepo "golang-project/repository/mongodb"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		dynamicCORSConfig := middleware.CORSConfig{
			AllowOrigins:     []string{"*", "*"},
			AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, "doktuz_apikey", "doktuz_apikey2"},
			AllowCredentials: true,
		}
		CORSMiddleware := middleware.CORSWithConfig(dynamicCORSConfig)
		CORSHandler := CORSMiddleware(next)
		return CORSHandler(c)
	}
}

func Cookies(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("sessionID")
		if err != nil {
			return security.Forbidden()
		}
		tknStr := cookie.Value
		if len(tknStr) > 0 {
			tkn, err := jwt.ParseWithClaims(tknStr, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte("secret"), nil
			})
			sessionsCollection := sessionRepo.GetSessionsCollection()
			if err != nil {
				return err
			}
			if !tkn.Valid {
				return security.Unauthorized()
			} else {
				fmt.Printf("\n\n\n\nlinea 46 middleware cookie passsssoooo %v\n\n", tknStr)
				sessionEnabled := sessionsCollection.ValidateSession(nil, tknStr)
				//session, err := sessionsCollection.GetSessionById(nil, tknStr)
				fmt.Printf("linea 45 middleware cookie %v\n\n", sessionEnabled)
				if sessionEnabled {
					return next(c)
				} else {
					return security.Unauthorized()
				}
			}
		} else {
			return security.SessionExpired()
		}
	}
}
