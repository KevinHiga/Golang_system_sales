package http

import (
	"context"
	"fmt"
	"golang-project/config/dbiface"
	"golang-project/config/security"
	"log"
	"net/http"

	_HttpDeliveryMiddleware "golang-project/config/middleware"
	"golang-project/models"
	userUcase "golang-project/usecase"

	"github.com/labstack/echo/v4"
)

func NewUsersHandler(g *echo.Group, ucol dbiface.CollectionAPI) {
	h := Handler{Col: ucol}
	g.POST("/login", h.login)
	g.POST("/logingoogle", h.LoginGoogle)
	g.POST("/forgot", h.ForgotPassword)
	g.Use(_HttpDeliveryMiddleware.Cookies)
	userGroup := g.Group("/user")
	userGroup.POST("/registrate", h.CreateUsersEndpoint)
	userGroup.POST("/logout", h.Logout)
	userGroup.GET("/all", h.GetUsersEndpoint)
	userGroup.GET("/one", h.User)
}

func (h *Handler) LoginGoogle(c echo.Context) error {
	var users []models.Users
	if err := c.Bind(&users); err != nil {
		log.Printf("Unable to bind : %v", err)
		return err
	}
	username, err := userUcase.LoginGoogle(context.Background(), users, c.Request().Body, h.Col, c)
	if err != nil {
		return security.CustomError(500, err.Error())
	}
	return c.JSONPretty(http.StatusOK, username, "  ")
}

func (h *Handler) login(c echo.Context) error {
	var users []models.Users
	if err := c.Bind(&users); err != nil {
		log.Printf("Unable to bind : %v", err)
		return err
	}
	username, err := userUcase.Login(context.Background(), users, c.Request().Body, h.Col, c)
	if err != nil {
		return security.CustomError(500, err.Error())
	}
	return c.JSONPretty(http.StatusOK, username, "  ")
}

func (h *Handler) ForgotPassword(c echo.Context) error {
	var users []models.Users
	if err := c.Bind(&users); err != nil {
		log.Printf("Unable to bind : %v", err)
		return err
	}
	username, err := userUcase.ForgotPassword(context.Background(), users, c.Request().Body, h.Col)
	if err != nil {
		return security.CustomError(500, err.Error())
	}
	return c.JSONPretty(http.StatusOK, username, "  ")
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
		return security.CustomError(500, err.Error())
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
		return security.Unauthorized()
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
