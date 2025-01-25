package model

type MockupStyles struct {
	Placement       string        `json:"placement" bson:"placement"`
	DisplayName     string        `json:"display_name" bson:"display_name"`
	Technique       string        `json:"technique" bson:"technique"`
	PrintAreaWidth  float64       `json:"print_area_width" bson:"print_area_width"`
	PrintAreaHeight float64       `json:"print_area_height" bson:"print_area_height"`
	PrintAreaType   string        `json:"print_area_type" bson:"print_area_type"`
	Dpi             int           `json:"dpi" bson:"dpi"`
	MockupStyles    []MockupStyle `json:"mockup_styles" bson:"mockup_styles"`
}

type MockupStyle struct {
	Id                   int    `json:"id" bson:"id"`
	CategoryName         string `json:"category_name" bson:"category_name"`
	ViewName             string `json:"view_name" bson:"view_name"`
	RestrictedToVariants []int  `json:"restricted_to_variants" bson:"restricted_to_variants"`
}
