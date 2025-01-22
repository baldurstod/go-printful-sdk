package model

type VariantPrice struct {
	Currency string           `json:"currency" bson:"currency" mapstructure:"currency"`
	Product  ProductPriceInfo `json:"product" bson:"product" mapstructure:"product"`
	Variant  VariantPriceInfo `json:"variant" bson:"variant" mapstructure:"variant"`
}

type ProductPriceInfo struct {
	ID         int                  `json:"id" bson:"id" mapstructure:"id"`
	Placements []PlacementPriceInfo `json:"placements" bson:"placements" mapstructure:"placements"`
}

type PlacementPriceInfo struct {
	ID              string            `json:"id" bson:"id" mapstructure:"id"`
	Price           string            `json:"price" bson:"price" mapstructure:"price"`
	DiscountedPrice string            `json:"discounted_price" bson:"discounted_price" mapstructure:"discounted_price"`
	Type            string            `json:"type" bson:"type" mapstructure:"type"`
	TechniqueKey    string            `json:"technique_key" bson:"technique_key" mapstructure:"technique_key"`
	Title           string            `json:"title" bson:"title" mapstructure:"title"`
	Options         []PlacementOption `json:"placement_options" bson:"placement_options" mapstructure:"placement_options"`
	Layers          []LayerPriceInfo  `json:"layers" bson:"layers" mapstructure:"layers"`
}

type LayerPriceInfo struct {
	AdditionalPrice string        `json:"additional_price" bson:"additional_price" mapstructure:"additional_price"`
	Type            string        `json:"type" bson:"type" mapstructure:"type"`
	Options         []LayerOption `json:"layer_options" bson:"layer_options" mapstructure:"layer_options"`
}

type VariantPriceInfo struct {
	ID         int                  `json:"id" bson:"id" mapstructure:"id"`
	Techniques []TechniquePriceInfo `json:"techniques" bson:"techniques" mapstructure:"techniques"`
}

type TechniquePriceInfo struct {
	Price           string `json:"price" bson:"price" mapstructure:"price"`
	DiscountedPrice string `json:"discounted_price" bson:"discounted_price" mapstructure:"discounted_price"`
	Key             string `json:"technique_key" bson:"technique_key" mapstructure:"technique_key"`
	DisplayName     string `json:"technique_display_name" bson:"technique_display_name" mapstructure:"technique_display_name"`
}
