package responses

import model "github.com/baldurstod/go-printful-sdk/model"

type VariantPricesResponse struct {
	Data  model.VariantPrice `json:"data" bson:"data" mapstructure:"data"`
	Links VariantLinks       `json:"_links" bson:"_links" mapstructure:"_links"`
}

type VariantLinks struct {
	Self           model.Link `json:"self" bson:"self" mapstructure:"self"`
	ProductDetails model.Link `json:"product_details" bson:"product_details" mapstructure:"product_details"`
	ProductPrices  model.Link `json:"product_prices" bson:"product_prices" mapstructure:"product_prices"`
}
