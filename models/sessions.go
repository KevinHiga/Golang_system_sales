package models

type Session struct {
	ID       string `json:"_id,omitempty" bson:"_id,omitempty"`
	Ssid string `json:"ssid" bson:"ssid"`
	Username     string `json:"username,omitempty" bson:"username,omitempty"`
	Enabled bool `json:"enabled,omitempty" bson:"enabled,omitempty"`
	Browser string `json:"browser,omitempty" bson:"browser,omitempty"`
	Os string `json:"os,omitempty" bson:"os,omitempty"`
}