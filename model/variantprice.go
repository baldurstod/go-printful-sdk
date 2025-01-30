package model

type VariantPrice struct {
	Currency string            `json:"currency" bson:"currency" mapstructure:"currency"`
	Product  ProductPriceInfo  `json:"product" bson:"product" mapstructure:"product"`
	Variant  VariantsPriceData `json:"variant" bson:"variant" mapstructure:"variant"`
}

type ProductPriceInfo struct {
	ID         int                    `json:"id" bson:"id" mapstructure:"id"`
	Placements []AdditionalPlacements `json:"placements" bson:"placements" mapstructure:"placements"`
}

type AdditionalPlacements struct {
	ID               string             `json:"id" bson:"id" mapstructure:"id"`
	Title            string             `json:"title" bson:"title" mapstructure:"title"`
	Type             string             `json:"type" bson:"type" mapstructure:"type"`
	TechniqueKey     string             `json:"technique_key" bson:"technique_key" mapstructure:"technique_key"`
	Price            string             `json:"price" bson:"price" mapstructure:"price"`
	DiscountedPrice  string             `json:"discounted_price" bson:"discounted_price" mapstructure:"discounted_price"`
	PlacementOptions []FileOptionPrices `json:"placement_options" bson:"placement_options" mapstructure:"placement_options"`
	Layers           []Layers           `json:"layers" bson:"layers" mapstructure:"layers"`
}

type Layers struct {
	Type            string              `json:"type" bson:"type" mapstructure:"type"`
	AdditionalPrice string              `json:"additional_price" bson:"additional_price" mapstructure:"additional_price"`
	Options         []LayerOptionPrices `json:"layer_options" bson:"layer_options" mapstructure:"layer_options"`
}

type LayerOptionPrices struct {
	Name        string            `json:"name" bson:"name"`
	Type        string            `json:"type" bson:"type"`
	Values      []any             `json:"values" bson:"values"`
	Description string            `json:"description" bson:"description"`
	Price       map[string]string `json:"price" bson:"price"`
}

type VariantsPriceData struct {
	ID         int                  `json:"id" bson:"id" mapstructure:"id"`
	Techniques []TechniquePriceInfo `json:"techniques" bson:"techniques" mapstructure:"techniques"`
}

type TechniquePriceInfo struct {
	Price           string `json:"price" bson:"price" mapstructure:"price"`
	DiscountedPrice string `json:"discounted_price" bson:"discounted_price" mapstructure:"discounted_price"`
	Key             string `json:"technique_key" bson:"technique_key" mapstructure:"technique_key"`
	DisplayName     string `json:"technique_display_name" bson:"technique_display_name" mapstructure:"technique_display_name"`
}

type FileOptionPrices struct {
	Name        string            `json:"name" bson:"name"`
	Type        string            `json:"type" bson:"type"`
	Values      []any             `json:"values" bson:"values"`
	Description string            `json:"description" bson:"description"`
	Price       map[string]string `json:"price" bson:"price"`
}
