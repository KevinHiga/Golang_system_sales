package models

type ProductType struct {
	ID          string `json:"_id,omitempty" bson:"_id,omitempty"`
	Name    string `json:"name" bson:"name"`
}