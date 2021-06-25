package models

type ChartsRender struct {
	ID      string  `json:"_id,omitempty" bson:"_id,omitempty"`
	X       int64   `json:"x" bson:"x"`
	Y       int64   `json:"y" bson:"y"`
	H       int64   `json:"h" bson:"h"`
	W       int64   `json:"w" bson:"w"`
	BaseURL string  `json:"baseUrl" bson:"baseUrl"`
	ChartID string  `json:"chartId" bson:"chartId"`
	Sesion  Session `json:"session,omitempty" bson:"session,omitempty"`
}
