package responses

import model "github.com/baldurstod/go-printful-sdk/model"

type CategoriesResponse struct {
	Data   []model.Category `json:"data" bson:"data" mapstructure:"data"`
	Paging `json:"paging" bson:"paging" mapstructure:"paging"`
	Links  `json:"_links" bson:"_links" mapstructure:"_links"`
}
