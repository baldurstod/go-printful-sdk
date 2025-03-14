package model

type Placement struct {
	Placement         string  `json:"placement" bson:"placement" mapstructure:"placement"`
	Technique         string  `json:"technique" bson:"technique" mapstructure:"technique"`
	PrintAreaType     string  `json:"print_area_type" bson:"print_area_type" mapstructure:"print_area_type"`
	Layers            []Layer `json:"layers" bson:"layers" mapstructure:"layers"`
	PlacementOptions  `json:"placement_options" bson:"placement_options" mapstructure:"placement_options"`
	Status            string `json:"status" bson:"status" mapstructure:"status"`
	StatusExplanation string `json:"status_explanation" bson:"status_explanation" mapstructure:"status_explanation"`
}

type PlacementOptions []PlacementOption

type PlacementOption struct {
	Name       string   `json:"name" bson:"name" mapstructure:"name"`
	Techniques []string `json:"techniques" bson:"techniques" mapstructure:"techniques"`
	Type       string   `json:"type" bson:"type" mapstructure:"type"`
	Values     any      `json:"values" bson:"values" mapstructure:"values"`
}

func NewPlacement() Placement {
	return Placement{
		PrintAreaType: "simple",
	}
}
