package model

type Placement struct {
	Placement string            `json:"placement" bson:"placement" mapstructure:"placement"`
	Technique string            `json:"technique" bson:"technique" mapstructure:"technique"`
	Layers    []Layer           `json:"layers" bson:"layers" mapstructure:"layers"`
	Options   []PlacementOption `json:"placement_options" bson:"placement_options" mapstructure:"placement_options"`
}

type PlacementOption struct {
	Name       string   `json:"name" bson:"name" mapstructure:"name"`
	Techniques []string `json:"techniques" bson:"techniques" mapstructure:"techniques"`
	Type       string   `json:"type" bson:"type" mapstructure:"type"`
	Values     any      `json:"values" bson:"values" mapstructure:"values"`
}
