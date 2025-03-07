package model

type ProductPrices struct {
	Currency string              `json:"currency" bson:"currency" mapstructure:"currency"`
	Product  ProductPriceInfo    `json:"product" bson:"product" mapstructure:"product"`
	Variants []VariantsPriceData `json:"variants" bson:"variants" mapstructure:"variants"`
}
