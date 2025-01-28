package model

type Placement struct {
	Placement         string  `json:"placement" bson:"placement"`
	Technique         string  `json:"technique" bson:"technique"`
	PrintAreaType     string  `json:"print_area_type" bson:"print_area_type"`
	Layers            []Layer `json:"layers" bson:"layers"`
	PlacementOptions  `json:"placement_options" bson:"placement_options"`
	Status            string `json:"status" bson:"status"`
	StatusExplanation string `json:"status_explanation" bson:"status_explanation"`
}

type PlacementOptions []PlacementOption

type PlacementOption struct {
	Name       string   `json:"name" bson:"name"`
	Techniques []string `json:"techniques" bson:"techniques"`
	Type       string   `json:"type" bson:"type"`
	Values     any      `json:"values" bson:"values"`
}

func NewPlacement() Placement {
	return Placement{
		PrintAreaType: "simple",
	}
}
