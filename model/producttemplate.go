package model

// TODO: adapt for api v2
type ProductTemplate struct {
	Version               int                         `json:"version" bson:"version"`
	MinDPI                int                         `json:"min_dpi" bson:"min_dpi"`
	VariantMapping        []TemplateVariantMapping    `json:"variant_mapping" bson:"variant_mapping"`
	Templates             []Template                  `json:"templates" bson:"templates"`
	ConflictingPlacements []TemplatePlacementConflict `json:"conflicting_placements" bson:"conflicting_placements"`
}

type TemplateVariantMapping struct {
	VariantID int                          `json:"variant_id" bson:"variant_id"`
	Templates []TemplateVariantMappingItem `json:"templates" bson:"templates"`
}
type TemplateVariantMappingItem struct {
	Placement  string `json:"placement" bson:"placement"`
	TemplateID int    `json:"template_id" bson:"template_id"`
}

type Template struct {
	TemplateID        int    `json:"template_id" bson:"template_id"`
	ImageURL          string `json:"image_url" bson:"image_url"`
	BackgroundURL     string `json:"background_url" bson:"background_url"`
	BackgroundColor   string `json:"background_color" bson:"background_color"`
	PrintfileID       int    `json:"printfile_id" bson:"printfile_id"`
	TemplateWidth     int    `json:"template_width" bson:"template_width"`
	TemplateHeight    int    `json:"template_height" bson:"template_height"`
	PrintAreaWidth    int    `json:"print_area_width" bson:"print_area_width"`
	PrintAreaHeight   int    `json:"print_area_height" bson:"print_area_height"`
	PrintAreaTop      int    `json:"print_area_top" bson:"print_area_top"`
	PrintAreaLeft     int    `json:"print_area_left" bson:"print_area_left"`
	IsTemplateOnFront bool   `json:"is_template_on_front" bson:"is_template_on_front"`
	Orientation       string `json:"orientation" bson:"orientation"`
}
type TemplatePlacementConflict struct {
	Placement string   `json:"placement" bson:"placement"`
	Conflicts []string `json:"conflicts" bson:"conflicts"`
}
