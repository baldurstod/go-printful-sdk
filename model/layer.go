package model

type Layer struct {
	Type    string        `json:"type" bson:"type" mapstructure:"type"`
	Options []LayerOption `json:"layer_options" bson:"layer_options" mapstructure:"layer_options"`
}

type LayerOption struct {
	Name       string   `json:"name" bson:"name" mapstructure:"name"`
	Techniques []string `json:"techniques" bson:"techniques" mapstructure:"techniques"`
	Type       string   `json:"type" bson:"type" mapstructure:"type"`
	Values     any      `json:"values" bson:"values" mapstructure:"values"`
}
