package usecase

import (
	"context"
	"golang-project/config/dbiface"
	"golang-project/models"
	chartRepo "golang-project/repository/mongodb"
	"io"

	"github.com/labstack/echo/v4"
)

func GetChartsRender(ctx context.Context, collection dbiface.CollectionAPI, c echo.Context) (*[]models.ChartsRender, error) {
	var charts []models.ChartsRender
	cookie, err := c.Cookie("sessionID")
	if err != nil {
		return nil, err
	}
	crender, err := chartRepo.GetChartsRender(ctx, collection)
	if err != nil {
		return nil, err
	}
	for _, chart := range *crender {
		chart.Sesion.Ssid = cookie.Value
		charts = append(charts, chart)
	}
	return &charts, nil
}

func ModifyChartsRender(ctx context.Context, id string, reqBody io.ReadCloser, collection dbiface.CollectionAPI) (models.ChartsRender, error) {
	return chartRepo.ModifyChartsRender(ctx, id, reqBody, collection)
}
