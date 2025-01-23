package model

// Note: Dimensions should be int as declared in openapi.json, but the actual API send floats
type MockupTemplates struct {
	CatalogVariantIDs   []int   `json:"catalog_variant_ids" bson:"catalog_variant_ids"`
	Placement           string  `json:"placement" bson:"placement"`
	Technique           string  `json:"technique" bson:"technique"`
	ImageURL            string  `json:"image_url" bson:"image_url"`
	BackgroundURL       string  `json:"background_url" bson:"background_url"`
	BackgroundColor     int     `json:"background_color" bson:"background_color"`
	PrintfileID         int     `json:"printfile_id" bson:"printfile_id"`
	TemplateWidth       float64 `json:"template_width" bson:"template_width"`
	TemplateHeight      float64 `json:"template_height" bson:"template_height"`
	PrintAreaWidth      float64 `json:"print_area_width" bson:"print_area_width"`
	PrintAreaHeight     float64 `json:"print_area_height" bson:"print_area_height"`
	PrintAreaTop        float64 `json:"print_area_top" bson:"print_area_top"`
	PrintAreaLeft       float64 `json:"print_area_left" bson:"print_area_left"`
	TemplatePositioning string  `json:"template_positioning" bson:"template_positioning"`
	Orientation         string  `json:"orientation" bson:"orientation"`
}
