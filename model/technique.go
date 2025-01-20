package model

type Technique struct {
	Key         string `json:"key" bson:"key" mapstructure:"key"`
	DisplayName string `json:"display_name" bson:"display_name" mapstructure:"display_name"`
	IsDefault   bool   `json:"is_default" bson:"is_default" mapstructure:"is_default"`
}
