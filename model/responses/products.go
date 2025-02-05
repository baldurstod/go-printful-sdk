package responses

import model "github.com/baldurstod/go-printful-sdk/model"

type ProductsResponse struct {
	Data   []model.Product `json:"data" bson:"data" mapstructure:"data"`
	Paging `json:"paging" bson:"paging" mapstructure:"paging"`
	Links  `json:"_links" bson:"_links" mapstructure:"_links"`
}

type ProductResponse struct {
	Data model.Product `json:"data" bson:"data" mapstructure:"data"`
}
