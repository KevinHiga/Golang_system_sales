package http

import (
	"context"
	"golang-project/config/dbiface"
	_HttpDeliveryMiddleware "golang-project/config/middleware"
	"net/http"

	chartUcase "golang-project/usecase"

	"github.com/labstack/echo/v4"
)

func NewCharsRenderHandler(g *echo.Group, crcol dbiface.CollectionAPI) {
	crh := Handler{Col: crcol}
	g.Use(_HttpDeliveryMiddleware.Cookies)
	g.GET("/all", crh.GetChartsRender)
	g.PUT("/modify/:id", crh.ModifyChartsRender)
}

func (h *Handler) GetChartsRender(c echo.Context) error {
	chartrender, err := chartUcase.GetChartsRender(context.Background(), h.Col, c)
	if err != nil {
		return err
	}
	return c.JSONPretty(http.StatusOK, chartrender, "  ")
}

func (h *Handler) ModifyChartsRender(c echo.Context) error {
	chartrender, err := chartUcase.ModifyChartsRender(context.Background(), c.Param("id"), c.Request().Body, h.Col)
	if err != nil {
		return err
	}
	return c.JSONPretty(http.StatusOK, chartrender, "  ")
}
