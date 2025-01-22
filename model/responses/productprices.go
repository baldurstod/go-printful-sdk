package responses

import model "github.com/baldurstod/go-printful-sdk/model"

type ProductPricesResponse struct {
	Data   model.ProductPrices `json:"data" bson:"data" mapstructure:"data"`
	Paging `json:"paging" bson:"paging" mapstructure:"paging"`
	Links  ProductPricesLinks `json:"_links" bson:"_links" mapstructure:"_links"`
}

type ProductPricesLinks struct {
	model.Links
	ProductDetails model.Link `json:"product_details" bson:"product_details" mapstructure:"product_details"`
}
