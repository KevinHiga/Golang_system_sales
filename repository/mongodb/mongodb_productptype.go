package mongodb

import (
	"context"
	"golang-project/config/dbiface"
	"golang-project/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func FindOrderProductsPType(ctx context.Context, collection dbiface.CollectionAPI) ([]models.Productptype, error) {
	var productptype []models.Productptype
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("Incapaz de encontrar el libro %+v", err)
	}
	err = cursor.All(ctx, &productptype)
	if err != nil {
		log.Printf("Incapaz de leer el cursor %+v", err)
	}
	return productptype, nil
}

func FindOrderPriceTier(ctx context.Context, collection dbiface.CollectionAPI) ([]models.Productptype, error) {
	var productptype []models.Productptype
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("Incapaz de encontrar el libro %+v", err)
	}
	err = cursor.All(ctx, &productptype)
	if err != nil {
		log.Printf("Incapaz de leer el cursor %+v", err)
	}
	return productptype, nil
}
