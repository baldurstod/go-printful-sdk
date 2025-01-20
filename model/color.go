package model

type Color struct {
	Name  string `json:"name" bson:"name" mapstructure:"name"`
	Value string `json:"value" bson:"value" mapstructure:"value"`
}
