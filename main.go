package main

import (
	"context"
	"fmt"
	config "golang-project/config/db"
	"log"

	_HttpDeliveryMiddleware "golang-project/config/middleware"
	_HttpDelivery "golang-project/delivery/http"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	c    *mongo.Client
	db   *mongo.Database
	ucol *mongo.Collection
	pcol *mongo.Collection
	ptcol *mongo.Collection
	cfg  config.Properties
)

func init() {
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatalf("La configuracion no puede ser leida: %v", err)
	}
	connectURI := fmt.Sprintf("%s", cfg.DBMongo)
	c, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectURI))
	if err != nil {
		log.Fatalf("Incapaz de conectarse a la base de datos: %v", err)
	}
	db = c.Database(cfg.DBName)
	ucol = db.Collection(cfg.UsersCollection)
	pcol = db.Collection(cfg.ProductCollection)
	ptcol = db.Collection(cfg.ProductTypeCollection)
}

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(_HttpDeliveryMiddleware.CORS)
	_HttpDelivery.NewUsersHandler(e, ucol)
	_HttpDelivery.NewProductHandler(e, pcol)
	_HttpDelivery.NewProductTypeHandler(e, ptcol)
	e.Logger.Infof("Escuchando en %s:%s", cfg.Host, cfg.Port)
	fmt.Println("Escuchando")
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Port)))
}
