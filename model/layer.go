package model

type Layer struct {
	Type           string `json:"type" bson:"type"`
	Url            string `json:"url" bson:"url"`
	*LayerOptions  `json:"layer_options" bson:"layer_options"`
	*LayerPosition `json:"position" bson:"position"`
}

type LayerOptions []LayerOption

type LayerOption struct {
	Name       string   `json:"name" bson:"name"`
	Techniques []string `json:"techniques" bson:"techniques"`
	Type       string   `json:"type" bson:"type"`
	Values     any      `json:"values" bson:"values"`
}

type LayerPosition struct {
	Width  float64 `json:"width" bson:"width"`
	Height float64 `json:"height" bson:"height"`
	Top    float64 `json:"top" bson:"top"`
	Left   float64 `json:"left" bson:"left"`
}
