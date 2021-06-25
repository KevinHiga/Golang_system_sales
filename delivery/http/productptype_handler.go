package http

import (
	"context"
	"golang-project/config/dbiface"
	_HttpDeliveryMiddleware "golang-project/config/middleware"
	"net/http"

	productTypeUcase "golang-project/usecase"

	"github.com/labstack/echo/v4"
)

func NewProductPTypeHandler(g *echo.Group, pptcol dbiface.CollectionAPI, pricol dbiface.CollectionAPI) {
	ppth := Handler{Col: pptcol}
	prih := Handler{Col: pricol}
	g.GET("/all", ppth.GetOrderbyProductsType)
	g.GET("/pricetier", prih.GetOrderbyPriceTier)
	g.Use(_HttpDeliveryMiddleware.Cookies)
}

func (h *Handler) GetOrderbyProductsType(c echo.Context) error {
	productsptype, err := productTypeUcase.FindOrderbyProductsType(context.Background(), h.Col)
	if err != nil {
		return err
	}
	return c.JSONPretty(http.StatusOK, productsptype, "  ")
}

func (h *Handler) GetOrderbyPriceTier(c echo.Context) error {
	productsptype, err := productTypeUcase.FindOrderbyPriceTier(context.Background(), h.Col)
	if err != nil {
		return err
	}
	return c.JSONPretty(http.StatusOK, productsptype, "  ")
}
