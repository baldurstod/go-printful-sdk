package responses

import model "github.com/baldurstod/go-printful-sdk/model"

type ProductsResponse struct {
	Data   []model.ProductInfo `json:"data" bson:"data" mapstructure:"data"`
	Paging `json:"paging" bson:"paging" mapstructure:"paging"`
}

type ProductResponse struct {
	Data model.ProductInfo `json:"data" bson:"data" mapstructure:"data"`
}
