package responses

import model "github.com/baldurstod/go-printful-sdk/model"

type ProductPricesResponse struct {
	Data   model.ProductPrices `json:"data" bson:"data" mapstructure:"data"`
	Paging `json:"paging" bson:"paging" mapstructure:"paging"`
}
