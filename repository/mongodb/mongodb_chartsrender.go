package mongodb

import (
	"context"
	"encoding/json"
	"golang-project/config/dbiface"
	"golang-project/models"
	"io"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func GetChartsRender(ctx context.Context, collection dbiface.CollectionAPI) (*[]models.ChartsRender, error) {
	var chartrender []models.ChartsRender
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("Incapaz de encontrar el libro %+v", err)
		return nil, err
	}
	err = cursor.All(ctx, &chartrender)
	if err != nil {
		log.Printf("Incapaz de leer el cursor %+v", err)
		return nil, err
	}
	return &chartrender, nil
}

func ModifyChartsRender(ctx context.Context, id string, reqBody io.ReadCloser, collection dbiface.CollectionAPI) (models.ChartsRender, error) {
	var chartrender models.ChartsRender
	//find if the product exits, if err return 404
	filter := bson.M{"_id": id}
	res := collection.FindOne(ctx, filter)
	log.Println(res)
	log.Println(filter)
	if err := res.Decode(&chartrender); err != nil {
		log.Printf("unable to decode to library :%v", err)
		return chartrender, err
	}

	if err := json.NewDecoder(reqBody).Decode(&chartrender); err != nil {
		log.Printf("unable to decode using reqbody : %v", err)
		return chartrender, err
	}

	//update the product, if err return 500
	_, err := collection.UpdateOne(ctx, filter, bson.M{"$set": chartrender})
	if err != nil {
		log.Printf("Incapaz de actualizar el libro : %v", err)
		return chartrender, err
	}
	return chartrender, nil
}
