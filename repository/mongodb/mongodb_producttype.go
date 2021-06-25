package mongodb

import (
	"context"
	"encoding/json"
	"golang-project/config/dbiface"
	"golang-project/models"
	"io"
	"log"

	"github.com/segmentio/ksuid"
	"go.mongodb.org/mongo-driver/bson"
)


func FindProductsType(ctx context.Context, collection dbiface.CollectionAPI) ([]models.ProductType, error) {
	var product []models.ProductType
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("Incapaz de encontrar el libro %+v", err)
	}
	err = cursor.All(ctx, &product)
	if err != nil {
		log.Printf("Incapaz de leer el cursor %+v", err)
	}
	return product, nil
}

func FindProductType(ctx context.Context, id string, collection dbiface.CollectionAPI) (models.ProductType, error) {
	var product models.ProductType
	res := collection.FindOne(ctx, bson.M{"_id": id})
	err := res.Decode(&product)
	if err != nil {
		return product, err
	}
	return product, nil
}

func ModifyProductType(ctx context.Context, id string, reqBody io.ReadCloser, collection dbiface.CollectionAPI) (models.ProductType, error) {
	var product models.ProductType
	//find if the product exits, if err return 404
	filter := bson.M{"_id": id}
	res := collection.FindOne(ctx, filter)
	log.Println(res)
	log.Println(filter)
	if err := res.Decode(&product); err != nil {
		log.Printf("unable to decode to library :%v", err)
		return product, err
	}

	if err := json.NewDecoder(reqBody).Decode(&product); err != nil {
		log.Printf("unable to decode using reqbody : %v", err)
		return product, err
	}

	//update the product, if err return 500
	_, err := collection.UpdateOne(ctx, filter, bson.M{"$set": product})
	if err != nil {
		log.Printf("Incapaz de actualizar el libro : %v", err)
		return product, err
	}
	return product, nil
}

func InsertProductType(ctx context.Context, products []models.Product, collection dbiface.CollectionAPI) ([]interface{}, error) {
	var insertedIds []interface{}
	for _, product := range products {
		product.ID = ksuid.New().String()
		insertID, err := collection.InsertOne(ctx, product)
		if err != nil {
			log.Printf("Incapaz de insertar en la base de datos:%v", err)
			return nil, err
		}
		insertedIds = append(insertedIds, insertID.InsertedID)
	}
	return insertedIds, nil
}

func DeleteProductType(ctx context.Context, id string, collection dbiface.CollectionAPI) (int64, error) {
	res, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		log.Printf("Incapaz de eliminar un producto : %v", err)
		return 0, err
	}
	return res.DeletedCount, nil
}