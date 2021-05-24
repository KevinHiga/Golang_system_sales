package models

type Product struct {
	ID          string `json:"_id,omitempty" bson:"_id,omitempty"`
	Name    string `json:"name" bson:"name"`
	Price        float64 `json:"price,omitempty" bson:"price,omitempty"`
	Stock    int64 `json:"stock,omitempty" bson:"stock,omitempty"`
	ProductTypeID    string `json:"producttypeid,omitempty" bson:"producttypeid,omitempty"`
}