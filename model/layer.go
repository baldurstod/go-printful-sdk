package model

type Layer struct {
	Type           string `json:"type" bson:"type" mapstructure:"type"`
	Url            string `json:"url" bson:"url" mapstructure:"url"`
	*LayerOptions  `json:"layer_options,omitempty" bson:"layer_options" mapstructure:"layer_options,omitempty"`
	*LayerPosition `json:"position,omitempty" bson:"position" mapstructure:"position,omitempty"`
}

type LayerOptions []LayerOption

type LayerOption struct {
	Name       string   `json:"name" bson:"name" mapstructure:"name"`
	Techniques []string `json:"techniques" bson:"techniques" mapstructure:"techniques"`
	Type       string   `json:"type" bson:"type" mapstructure:"type"`
	Values     any      `json:"values" bson:"values" mapstructure:"values"`
}

type LayerPosition struct {
	Width  float64 `json:"width" bson:"width" mapstructure:"width"`
	Height float64 `json:"height" bson:"height" mapstructure:"height"`
	Top    float64 `json:"top" bson:"top" mapstructure:"top"`
	Left   float64 `json:"left" bson:"left" mapstructure:"left"`
}
