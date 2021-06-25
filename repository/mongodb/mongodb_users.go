package mongodb

import (
	"context"
	"fmt"
	"golang-project/config/dbiface"
	"golang-project/config/security"
	"golang-project/models"
	"io"
	"log"

	"github.com/segmentio/ksuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InsertUser(ctx context.Context, users []models.Users, collection dbiface.CollectionAPI) ([]interface{}, error) {
	var insertedIds []interface{}
	for _, user := range users {
		user.ID = ksuid.New().String()
		exists, err := FindUserName(ctx, user.UserName, collection)
		if exists != nil {
			err := fmt.Errorf("The username already exists")
			return nil, err
		}
		exists, err = FindMail(ctx, user.Mail, collection)
		if exists != nil {
			err := fmt.Errorf("The mail already exists")
			return nil, err
		}
		user.Password, err = security.EncryptPassword(user.Password)
		if err != nil {
			log.Printf("Incapaz de insertar en la base de datos:%v", err)
			return nil, err
		}
		insertID, err := collection.InsertOne(ctx, user)
		if err != nil {
			log.Printf("Incapaz de insertar en la base de datos:%v", err)
			return nil, err
		}
		insertedIds = append(insertedIds, insertID.InsertedID)
	}
	return insertedIds, nil
}

func FindUsers(ctx context.Context, collection dbiface.CollectionAPI) (*[]models.Users, error) {
	var user []models.Users
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("Incapaz de encontrar el libro %+v", err)
		return nil, err
	}
	err = cursor.All(ctx, &user)
	if err != nil {
		log.Printf("Incapaz de leer el cursor %+v", err)
		return nil, err
	}
	return &user, nil
}

func FindUser(ctx context.Context, id string, collection dbiface.CollectionAPI) (models.Users, error) {
	var user models.Users
	res := collection.FindOne(ctx, bson.M{"_id": id})
	err := res.Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func FindUserName(ctx context.Context, username string, collection dbiface.CollectionAPI) (*models.Users, error) {
	var user models.Users
	res := collection.FindOne(ctx, bson.M{"username": username})
	if err := res.Decode(&user); err != nil {
		log.Printf("unable to decode to user %+v", err)
		return nil, err
	}
	return &user, nil
}

func FindMail(ctx context.Context, mail string, collection dbiface.CollectionAPI) (*models.Users, error) {
	var user models.Users
	res := collection.FindOne(ctx, bson.M{"mail": mail})
	if err := res.Decode(&user); err != nil {
		log.Printf("unable to decode to user %+v", err)
		return nil, err
	}
	return &user, nil
}

func Login(ctx context.Context, inputs []models.Users, reqBody io.ReadCloser, collection dbiface.CollectionAPI) (*[]models.Users, error) {
	for _, input := range inputs {
		exists, err := FindUserName(ctx, input.UserName, collection)
		if exists == nil {
			err = fmt.Errorf("The username does not exist")
			return nil, err
		}
		/*
			exists, err = FindMail(ctx, user.Mail, collection)
			if exists != nil {
				log.Println("El mail ya existe")
				err := fmt.Errorf("El mail ya existe")
				return nil, err
			}
		*/
		input.Password, err = security.EncryptPassword(input.Password)
		if err != nil {
			log.Printf("Incapaz de insertar en la base de datos:%v", err)
			return nil, err
		}
		opts := options.Update().SetUpsert(true)
		filter := bson.M{"username": input.UserName}
		update := bson.D{{"$set", bson.D{{"password", input.Password}}}}
		_, err = collection.UpdateOne(ctx, filter, update, opts)
		if err != nil {
			log.Printf("Incapaz de actualizar el libro : %v", err)
			return nil, err
		}
	}
	return &inputs, nil
}

func ForgotPassword(ctx context.Context, inputs []models.Users, reqBody io.ReadCloser, collection dbiface.CollectionAPI) (*[]models.Users, error) {
	for _, input := range inputs {
		exists, err := FindUserName(ctx, input.UserName, collection)
		if exists == nil {
			err = fmt.Errorf("The username does not exist")
			return nil, err
		}
		/*
			exists, err = FindMail(ctx, user.Mail, collection)
			if exists != nil {
				log.Println("El mail ya existe")
				err := fmt.Errorf("El mail ya existe")
				return nil, err
			}
		*/
		input.Password, err = security.EncryptPassword(input.Password)
		if err != nil {
			log.Printf("Incapaz de insertar en la base de datos:%v", err)
			return nil, err
		}
		opts := options.Update().SetUpsert(true)
		filter := bson.M{"username": input.UserName}
		update := bson.D{{"$set", bson.D{{"password", input.Password}}}}
		_, err = collection.UpdateOne(ctx, filter, update, opts)
		if err != nil {
			log.Printf("Incapaz de actualizar el libro : %v", err)
			return nil, err
		}
	}
	return &inputs, nil
}
