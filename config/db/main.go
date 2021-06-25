package config

import (
	"context"
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	C      *mongo.Client
	DB     *mongo.Database
	Ucol   *mongo.Collection
	Pcol   *mongo.Collection
	Ptcol  *mongo.Collection
	Pptcol *mongo.Collection
	Pricol *mongo.Collection
	Crcol  *mongo.Collection
	Cfg    Properties
)

func init() {
	if err := cleanenv.ReadEnv(&Cfg); err != nil {
		log.Fatalf("La configuracion no puede ser leida: %v", err)
	}
	connectURI := fmt.Sprintf("%s", Cfg.DBMongo)
	c, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectURI))
	if err != nil {
		log.Fatalf("Incapaz de conectarse a la base de datos: %v", err)
	}
	DB = c.Database(Cfg.DBName)
	Ucol = DB.Collection(Cfg.UsersCollection)
	Pcol = DB.Collection(Cfg.ProductCollection)
	Ptcol = DB.Collection(Cfg.ProductTypeCollection)
	Pptcol = DB.Collection(Cfg.ProductPTypeCollection)
	Pricol = DB.Collection(Cfg.ProductPriceTierCollection)
	Crcol = DB.Collection(Cfg.ChartsRenderCollection)
}
