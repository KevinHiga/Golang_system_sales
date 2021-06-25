package models

type Productptype struct {
	ID          string `json:"_id,omitempty" bson:"_id,omitempty"`
	Name    string `json:"names" bson:"names"`
	Price        float64 `json:"price,omitempty" bson:"price,omitempty"`
	Stock    int64 `json:"stock,omitempty" bson:"stock,omitempty"`
	Item    int64 `json:"item" bson:"item"`
	ProductpType    ProductType `json:"producttype,omitempty" bson:"producttype,omitempty"`
}