package http

import (
	"context"
	"fmt"
	"golang-project/config/dbiface"
	_HttpDeliveryMiddleware "golang-project/config/middleware"
	"log"
	"net/http"

	productUcase "golang-project/usecase"

	"golang-project/models"

	"github.com/labstack/echo/v4"
)

func NewProductHandler(g *echo.Group, pcol dbiface.CollectionAPI) {
	ph := Handler{Col: pcol}
	g.Use(_HttpDeliveryMiddleware.Cookies)
	g.GET("/all", ph.GetProductsEndpoint)
	g.GET("/:id", ph.GetProductEndpoint)
	g.DELETE("/delete/:id", ph.DeleteProductEndpoint)
	g.PUT("/:id", ph.UpdateProductEndpoint)
	g.POST("/create", ph.CreateProductEndpoint)
}

func (h *Handler) GetProductsEndpoint(c echo.Context) error {
	products, err := productUcase.FindProductAllData(context.Background(), h.Col)
	if err != nil {
		return err
	}
	return c.JSONPretty(http.StatusOK, products, "  ")
}

func (h *Handler) GetProductEndpoint(c echo.Context) error {
	product, err := productUcase.FindProductOneData(context.Background(), c.Param("id"), h.Col)
	if err != nil {
		return err
	}
	return c.JSONPretty(http.StatusOK, product, "  ")
}

func (h *Handler) UpdateProductEndpoint(c echo.Context) error {
	product, err := productUcase.UpdateProductData(context.Background(), c.Param("id"), c.Request().Body, h.Col)
	if err != nil {
		return err
	}
	return c.JSONPretty(http.StatusOK, product, "  ")
}

func (h *Handler) CreateProductEndpoint(c echo.Context) error {
	var products []models.Product
	fmt.Println(products)
	if err := c.Bind(&products); err != nil {
		log.Printf("Unable to bind : %v", err)
		return err
	}
	IDs, err := productUcase.CreateProductData(context.Background(), products, h.Col)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, IDs)
}

func (h *Handler) DeleteProductEndpoint(c echo.Context) error {
	log.Println("paso\n\n")
	delCount, err := productUcase.DeleteProductData(context.Background(), c.Param("id"), h.Col)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, delCount)
}
