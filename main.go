package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	config "golang-project/config/db"
	_HttpDeliveryMiddleware "golang-project/config/middleware"
	_HttpDelivery "golang-project/delivery/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func uploadFile(c echo.Context) error {
	fmt.Println("File Upload Endpoint Hit")

	//r.ParseMultipartForm(10 << 20)
	c.Request().ParseMultipartForm(10 << 20)
	file, err := c.FormFile("myFile")
	//file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	fmt.Printf("Uploaded File: %+v\n", file.Filename)
	fmt.Printf("File Size: %+v\n", file.Size)
	fmt.Printf("MIME Header: %+v\n", file.Header)
	tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()
	fileBytes, err := ioutil.ReadAll(src)
	if err != nil {
		fmt.Println(err)
	}
	tempFile.Write(fileBytes)
	return c.JSONPretty(http.StatusOK, fileBytes, "  ")
}

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(_HttpDeliveryMiddleware.CORS)
	apiGroup := e.Group("/api")
	productGroup := apiGroup.Group("/product")
	chartGroup := apiGroup.Group("/chart")
	producTypeGroup := productGroup.Group("/type")
	producPTypeGroup := productGroup.Group("/ptype")
	_HttpDelivery.NewProductHandler(productGroup, config.Pcol)
	_HttpDelivery.NewUsersHandler(apiGroup, config.Ucol)
	_HttpDelivery.NewProductTypeHandler(producTypeGroup, config.Ptcol)
	_HttpDelivery.NewProductPTypeHandler(producPTypeGroup, config.Pptcol, config.Pricol)
	_HttpDelivery.NewCharsRenderHandler(chartGroup, config.Crcol)
	e.POST("/upload", uploadFile)
	e.GET("/api/helloworld", helloworld)
	e.Logger.Infof("Escuchando en %s:%s", config.Cfg.Host, config.Cfg.Port)
	fmt.Println("Escuchando")
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.Cfg.Port)))
}

func helloworld(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, "hello world", "  ")
}
