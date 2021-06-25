package security

import (
	"golang-project/models"

	"github.com/labstack/echo/v4"
)

func Forbidden() *echo.HTTPError {
	json := models.RequestErrors{Code: 403, Message: "Forbidden"}
	return echo.NewHTTPError(json.Code, json)
}

func Unauthorized() *echo.HTTPError {
	json := models.RequestErrors{Code: 401, Message: "Unauthorized"}
	return echo.NewHTTPError(json.Code, json)
}

func SessionExpired() *echo.HTTPError {
	json := models.RequestErrors{Code: 440, Message: "Session Expired"}
	return echo.NewHTTPError(json.Code, json)
}

func ErrorParsingToken() *echo.HTTPError {
	json := models.RequestErrors{Code: 500, Message: "Error parsing jwt"}
	return echo.NewHTTPError(json.Code, json)
}

func ErrorDatabase() *echo.HTTPError {
	json := models.RequestErrors{Code: 500, Message: "Error database"}
	return echo.NewHTTPError(json.Code, json)
}

func DocumentAlreadyExists() *echo.HTTPError {
	json := models.RequestErrors{Code: 409, Message: "Document already exists"}
	return echo.NewHTTPError(json.Code, json)
}

func DocumentsNotDeleted() *echo.HTTPError {
	json := models.RequestErrors{Code: 500, Message: "Documents was not deleted"}
	return echo.NewHTTPError(json.Code, json)
}

func CustomError(code int, message string) *echo.HTTPError {
	json := models.RequestErrors{Code: code, Message: message}
	return echo.NewHTTPError(json.Code, json)
}

func UnprocessableEntity() *echo.HTTPError {
	json := models.RequestErrors{Code: 422, Message: "Invalid request body"}
	return echo.NewHTTPError(json.Code, json)
}