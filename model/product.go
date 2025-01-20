package model

type Product struct {
	ID             int             `json:"id" bson:"id" mapstructure:"id"`
	MainCategoryID int             `json:"main_category_id" bson:"main_category_id" mapstructure:"main_category_id"`
	Type           string          `json:"type" bson:"type" mapstructure:"type"`
	Name           string          `json:"name" bson:"name" mapstructure:"name"`
	Brand          string          `json:"brand" bson:"brand" mapstructure:"brand"`
	Model          string          `json:"model" bson:"model" mapstructure:"model"`
	Image          string          `json:"image" bson:"image" mapstructure:"image"`
	VariantCount   int             `json:"variant_count" bson:"variant_count" mapstructure:"variant_count"`
	IsDiscontinued bool            `json:"is_discontinued" bson:"is_discontinued" mapstructure:"is_discontinued"`
	Description    string          `json:"description" bson:"description" mapstructure:"description"`
	Sizes          []string        `json:"sizes" bson:"sizes" mapstructure:"sizes"`
	Colors         []Color         `json:"colors" bson:"colors" mapstructure:"colors"`
	Techniques     []Technique     `json:"techniques" bson:"techniques" mapstructure:"techniques"`
	Placements     []Placement     `json:"placements" bson:"placements" mapstructure:"placements"`
	Options        []ProductOption `json:"product_options" bson:"product_options" mapstructure:"product_options"`
	Links          ProductLinks    `json:"_links" bson:"_links" mapstructure:"_links"`
}

// TODO: merge with PlacementOption and LayerOption ?
type ProductOption struct {
	Name       string   `json:"name" bson:"name" mapstructure:"name"`
	Techniques []string `json:"techniques" bson:"techniques" mapstructure:"techniques"`
	Type       string   `json:"type" bson:"type" mapstructure:"type"`
	Values     any      `json:"values" bson:"values" mapstructure:"values"`
}

type ProductLinks struct {
	Self                Link `json:"self" bson:"self" mapstructure:"self"`
	Categories          Link `json:"categories" bson:"categories" mapstructure:"categories"`
	ProductAvailability Link `json:"product_availability" bson:"product_availability" mapstructure:"product_availability"`
	ProductImages       Link `json:"product_images" bson:"product_images" mapstructure:"product_images"`
	ProductPrices       Link `json:"product_prices" bson:"product_prices" mapstructure:"product_prices"`
	ProductSizes        Link `json:"product_sizes" bson:"product_sizes" mapstructure:"product_sizes"`
	Variants            Link `json:"variants" bson:"variants" mapstructure:"variants"`
}
