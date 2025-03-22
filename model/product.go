package model

type ProductInfo struct {
	Product `bson:"inline" mapstructure:",squash"`
	Links   ProductLinks `json:"_links" bson:"_links" mapstructure:"_links"`
}

type Product struct {
	ID                int                `json:"id" bson:"id" mapstructure:"id"`
	MainCategoryID    int                `json:"main_category_id" bson:"main_category_id" mapstructure:"main_category_id"`
	Type              string             `json:"type" bson:"type" mapstructure:"type"`
	Name              string             `json:"name" bson:"name" mapstructure:"name"`
	Brand             string             `json:"brand" bson:"brand" mapstructure:"brand"`
	Model             string             `json:"model" bson:"model" mapstructure:"model"`
	Image             string             `json:"image" bson:"image" mapstructure:"image"`
	VariantCount      int                `json:"variant_count" bson:"variant_count" mapstructure:"variant_count"`
	CatalogVariantIDs []int              `json:"catalog_variant_ids" bson:"catalog_variant_ids"`
	IsDiscontinued    bool               `json:"is_discontinued" bson:"is_discontinued" mapstructure:"is_discontinued"`
	Description       string             `json:"description" bson:"description" mapstructure:"description"`
	Sizes             []string           `json:"sizes" bson:"sizes" mapstructure:"sizes"`
	Colors            []Color            `json:"colors" bson:"colors" mapstructure:"colors"`
	Techniques        []Technique        `json:"techniques" bson:"techniques" mapstructure:"techniques"`
	Placements        []ProductPlacement `json:"placements" bson:"placements" mapstructure:"placements"`
	ProductOptions    []CatalogOption    `json:"product_options" bson:"product_options" mapstructure:"product_options"`
}

type ProductPlacement struct {
	DesignPlacement       `bson:"inline"`
	ConflictingPlacements []string `json:"conflicting_placements" bson:"conflicting_placements"`
}

type CatalogOption struct {
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
