package models

import "github.com/dgrijalva/jwt-go"

type Users struct {
	ID       string `json:"_id,omitempty" bson:"_id,omitempty"`
	UserName string `json:"username" bson:"username"`
	Mail     string `json:"mail,omitempty" bson:"mail,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
	Admin bool `json:"admin,omitempty" bson:"admin,omitempty"`
	Sesion    Session `json:"Session,omitempty" bson:"Session,omitempty"`
	jwt.StandardClaims
}