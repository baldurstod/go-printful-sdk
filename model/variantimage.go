package model

type VariantImages struct {
	CatalogVariantId  int     `json:"catalog_variant_id" mapstructure:"catalog_variant_id"`
	Color             string  `json:"color" mapstructure:"color"`
	PrimaryHexColor   string  `json:"primary_hex_color" mapstructure:"primary_hex_color"`
	SecondaryHexColor string  `json:"secondary_hex_color" mapstructure:"secondary_hex_color"`
	Images            []Image `json:"images" mapstructure:"images"`
}

type Image struct {
	Placement       string `json:"placement" mapstructure:"placement"`
	ImageUrl        string `json:"image_url" mapstructure:"image_url"`
	BackgroundColor string `json:"background_color" mapstructure:"background_color"`
	BackgroundImage string `json:"background_image" mapstructure:"background_image"`
	MockupStyleId   int    `json:"mockup_style_id" mapstructure:"mockup_style_id"`
}
