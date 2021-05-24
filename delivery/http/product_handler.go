package http

import (
	"context"
	"fmt"
	"golang-project/config/dbiface"
	"log"
	"net/http"

	productUcase "golang-project/usecase"

	"golang-project/models"

	"github.com/labstack/echo/v4"
)

func NewProductHandler(e *echo.Echo, pcol dbiface.CollectionAPI) {
	ph := Handler{Col: pcol}
	e.GET("/products", ph.GetProductsEndpoint)
	e.GET("/product/:id", ph.GetProductEndpoint)
	e.DELETE("/product/:id", ph.DeleteProductEndpoint)
	e.PUT("/product/:id", ph.UpdateProductEndpoint)
	e.POST("/product", ph.CreateProductEndpoint)
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
	delCount, err := productUcase.DeleteProductData(context.Background(), c.Param("id"), h.Col)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, delCount)
}
