package responses

import model "github.com/baldurstod/go-printful-sdk/model"

type CountriesResponse struct {
	Data   []model.Country `json:"data" bson:"data" mapstructure:"data"`
	Paging `json:"paging" bson:"paging" mapstructure:"paging"`
	Links  `json:"_links" bson:"_links" mapstructure:"_links"`
}
