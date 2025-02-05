package model

type DesignPlacement struct {
	Placement        string          `json:"placement" bson:"placement"`
	Technique        string          `json:"technique" bson:"technique"`
	PrintAreaWidth   float64         `json:"print_area_width" bson:"print_area_width"`
	PrintAreaHeight  float64         `json:"print_area_height" bson:"print_area_height"`
	Layers           []FileLayer     `json:"layers" bson:"layers"`
	PlacementOptions []CatalogOption `json:"placement_options" bson:"placement_options"`
}
