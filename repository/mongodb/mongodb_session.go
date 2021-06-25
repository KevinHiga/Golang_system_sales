package mongodb

import (
	"context"
	"fmt"
	config "golang-project/config/db"
	"golang-project/config/dbiface"
	"golang-project/models"
	"log"

	"github.com/segmentio/ksuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SessionsCollection struct {
	Ca *mongo.Collection
}

func GetSessionsCollection() *SessionsCollection {
	collections := config.DB.Collection(config.Cfg.SessionCollection)
	return &SessionsCollection{Ca: collections}
}

func GetSessions(ctx context.Context, collection dbiface.CollectionAPI) ([]models.Session, error) {
	var session []models.Session
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("Incapaz de encontrar el libro %+v", err)
	}
	err = cursor.All(ctx, &session)
	if err != nil {
		log.Printf("Incapaz de leer el cursor %+v", err)
	}
	return session, nil
}

func (s *SessionsCollection) GetSessionById(ctx context.Context, ssid string) (*models.Session, error) {
	var session models.Session
	res := s.Ca.FindOne(ctx, bson.M{"_id": ssid})
	err := res.Decode(&session)
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func (s *SessionsCollection) AddSession(ctx context.Context, session *models.Session) error {
	if len(session.ID) == 0 {
		session.ID = ksuid.New().String()
	}
	_, err := s.Ca.InsertOne(ctx, session)
	if err != nil {
		log.Printf("Incapaz de insertar en la base de datos:%v", err)
		return err
	}
	return nil
}

func (s *SessionsCollection) ValidateSession(ctx context.Context, ssid string) bool {
	var session models.Session
	res := s.Ca.FindOne(ctx, bson.M{"ssid": ssid, "enabled": true})
	fmt.Printf("\n\n\nlinea 81 session %v\n\n", session)
	err := res.Decode(&session)
	fmt.Printf("\n\n\nlinea 81 session %v\n\n", res.Decode(&session))
	if err != nil {
		return false
	}
	return true
}
