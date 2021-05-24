package http

import (
	"context"
	"fmt"
	"golang-project/config/dbiface"
	"golang-project/models"
	"log"
	"net/http"

	productTypeUcase "golang-project/usecase"

	"github.com/labstack/echo/v4"
)

func NewProductTypeHandler(e *echo.Echo, pcol dbiface.CollectionAPI) {
	pth := Handler{Col: pcol}
	e.GET("/products/type", pth.GetProductsTypeEndpoint)
	e.GET("/product/type/:id", pth.GetProductTypeEndpoint)
	e.DELETE("/product/type/:id", pth.DeleteProductTypeEndpoint)
	e.PUT("/product/type/:id", pth.UpdateProductTypeEndpoint)
	e.POST("/product/type", pth.CreateProductTypeEndpoint)
}

func (h *Handler) GetProductsTypeEndpoint(c echo.Context) error {
	productstype, err := productTypeUcase.FindProductTypeAllData(context.Background(), h.Col)
	if err != nil {
		return err
	}
	return c.JSONPretty(http.StatusOK, productstype, "  ")
}

func (h *Handler) GetProductTypeEndpoint(c echo.Context) error {
	producttype, err := productTypeUcase.FindProductTypeOneData(context.Background(), c.Param("id"), h.Col)
	if err != nil {
		return err
	}
	return c.JSONPretty(http.StatusOK, producttype, "  ")
}

func (h *Handler) UpdateProductTypeEndpoint(c echo.Context) error {
	producttype, err := productTypeUcase.UpdateProductTypeData(context.Background(), c.Param("id"), c.Request().Body, h.Col)
	if err != nil {
		return err
	}
	return c.JSONPretty(http.StatusOK, producttype, "  ")
}

func (h *Handler) CreateProductTypeEndpoint(c echo.Context) error {
	var productstype []models.Product
	fmt.Println(productstype)
	if err := c.Bind(&productstype); err != nil {
		log.Printf("Unable to bind : %v", err)
		return err
	}
	IDs, err := productTypeUcase.CreateProductTypeData(context.Background(), productstype, h.Col)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, IDs)
}

func (h *Handler) DeleteProductTypeEndpoint(c echo.Context) error {
	delCount, err := productTypeUcase.DeleteProductTypeData(context.Background(), c.Param("id"), h.Col)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, delCount)
}