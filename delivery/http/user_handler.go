package http

import (
	"context"
	"fmt"
	"golang-project/config/dbiface"
	"log"
	"net/http"

	"golang-project/models"
	userUcase "golang-project/usecase"

	"github.com/labstack/echo/v4"
)

func NewUsersHandler(e *echo.Echo, ucol dbiface.CollectionAPI) {
	h := Handler{Col: ucol}
	e.POST("/user/registrate", h.CreateUsersEndpoint)
	e.POST("/user/login", h.Login)
	e.POST("/user/logout", h.Logout)
	e.POST("/user/forgot", h.ForgotPassword)
	e.GET("/users", h.GetUsersEndpoint)
	e.GET("/user", h.User)
}

func (h *Handler) CreateUsersEndpoint(c echo.Context) error {
	var users []models.Users
	fmt.Println(users)
	if err := c.Bind(&users); err != nil {
		log.Printf("Unable to bind : %v", err)
		return err
	}
	IDs, err := userUcase.CreateUsersData(context.Background(), users, h.Col)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, IDs)
}

func (h *Handler) GetUsersEndpoint(c echo.Context) error {
	user, err := userUcase.FindUserAllData(context.Background(), h.Col)
	if err != nil {
        c.Error(err)
		log.Printf("Unable to bind : %v", err)
		return c.JSONPretty(http.StatusInternalServerError, err.Error(), "  ")
	}
	return c.JSONPretty(http.StatusOK, user, "  ")
}

func (h *Handler) User(c echo.Context) error {
	user, err := userUcase.User(context.Background(), h.Col, c)
	if err != nil {
		return c.JSONPretty(http.StatusUnauthorized, user, "  ")
	}
	return c.JSONPretty(http.StatusOK, user, "  ")
}

func (h *Handler) Logout(c echo.Context) error {
	err := userUcase.Logout(c)
	if err != nil {
		return err
	}
	return c.JSONPretty(http.StatusOK, nil, "  ")
}

func (h *Handler) Login(c echo.Context) error {
	var users []models.Users
	if err := c.Bind(&users); err != nil {
		log.Printf("Unable to bind : %v", err)
		return err
	}
	username, err := userUcase.Login(context.Background(), users, h.Col, c)
	if err != nil {
		if err.Error() == "Contraseña equivocada" {
			
		}
		return c.JSONPretty(http.StatusUnauthorized, err.Error(), "  ")
	}
	return c.JSONPretty(http.StatusOK, username, "  ")
}

func (h *Handler) ForgotPassword(c echo.Context) error {
	var users models.Users
	if err := c.Bind(&users); err != nil {
		log.Printf("Unable to bind : %v", err)
		return err
	}
	username, err := userUcase.ForgotPassword(context.Background(), users, c.Request().Body, h.Col)
	if err != nil {
		if err.Error() == "Contraseña equivocada" {
			
		}
		return c.JSONPretty(http.StatusUnauthorized, err.Error(), "  ")
	}
	return c.JSONPretty(http.StatusOK, username, "  ")
}
