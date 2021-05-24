package mongodb

import (
	"context"
	"encoding/json"
	"fmt"
	"golang-project/config/dbiface"
	"golang-project/config/security"
	"golang-project/models"
	"io"
	"log"

	"github.com/segmentio/ksuid"
	"go.mongodb.org/mongo-driver/bson"
)

func InsertUser(ctx context.Context, users []models.Users, collection dbiface.CollectionAPI) ([]interface{}, error) {
	var insertedIds []interface{}
	for _, user := range users {
		user.ID = ksuid.New().String()
		//err := fmt.Errorf("")
		exists, err := FindUserName(ctx, user.UserName, collection)
		if exists != nil {
			log.Println("El username ya existe")
			err := fmt.Errorf("El username ya existe")
			return nil, err
		}
		exists, err = FindMail(ctx, user.Mail, collection)
		if exists != nil {
			log.Println("El mail ya existe")
			err := fmt.Errorf("El mail ya existe")
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

func ForgotPassword(ctx context.Context, users models.Users, reqBody io.ReadCloser, collection dbiface.CollectionAPI) (*models.Users, error) {
	exists, err := FindUserName(ctx, users.UserName, collection)
	if exists == nil {
		log.Println("El username no existe")
		err := fmt.Errorf("El username no existe")
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
	users.Password, err = security.EncryptPassword(users.Password)
	if err != nil {
		log.Printf("Incapaz de insertar en la base de datos:%v", err)
		return nil, err
	}
	filter := bson.M{"username": users.UserName}
	res := collection.FindOne(ctx, filter)
	if err := res.Decode(&users); err != nil {
		log.Printf("unable to decode to library :%v", err)
		return &users, err
	}
	if err := json.NewDecoder(reqBody).Decode(&users); err != nil {
		log.Printf("unable to decode using reqbody : %v", err)
		return &users, err
	}
	_, err = collection.UpdateOne(ctx, filter, bson.M{"$set": users.Password})
	if err != nil {
		log.Printf("Incapaz de actualizar el libro : %v", err)
		return &users, err
	}
	return &users, nil
}
