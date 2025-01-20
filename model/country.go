package model

type Country struct {
	Name   string  `json:"name" bson:"name" mapstructure:"name"`
	Code   string  `json:"code" bson:"code" mapstructure:"code"`
	Region string  `json:"region" bson:"region" mapstructure:"region"`
	States []State `json:"states" bson:"states" mapstructure:"states"`
}

type State struct {
	Name string `json:"name" bson:"name" mapstructure:"name"`
	Code string `json:"code" bson:"code" mapstructure:"code"`
}
