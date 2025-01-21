package model

type Variant struct {
	ID               int            `json:"id" bson:"id" mapstructure:"id"`
	Name             string         `json:"name" bson:"name" mapstructure:"name"`
	CatalogProductID int            `json:"catalog_product_id" bson:"catalog_product_id" mapstructure:"catalog_product_id"`
	Color            string         `json:"color" bson:"color" mapstructure:"color"`
	ColorCode        string         `json:"color_code" bson:"color_code" mapstructure:"color_code"`
	ColorCode2       string         `json:"color_code2" bson:"color_code2" mapstructure:"color_code2"`
	Image            string         `json:"image" bson:"image" mapstructure:"image"`
	Size             string         `json:"size" bson:"size" mapstructure:"size"`
	Availability     []Availability `json:"availability" bson:"availability" mapstructure:"availability"`
	Links            VariantLinks   `json:"_links" bson:"_links" mapstructure:"_links"`
}

type VariantLinks struct {
	Self                Link `json:"self" bson:"self" mapstructure:"self"`
	ProductDetails      Link `json:"product_details" bson:"product_details" mapstructure:"product_details"`
	ProductVariants     Link `json:"product_variants" bson:"product_variants" mapstructure:"product_variants"`
	VariantAvailability Link `json:"variant_availability" bson:"variant_availability" mapstructure:"variant_availability"`
	VariantImages       Link `json:"variant_images" bson:"variant_images" mapstructure:"variant_images"`
	VariantPrices       Link `json:"variant_prices" bson:"variant_prices" mapstructure:"variant_prices"`
}

type Availability struct {
	Region string `json:"region" bson:"region" mapstructure:"region"`
	Status string `json:"status" bson:"status" mapstructure:"status"`
}
