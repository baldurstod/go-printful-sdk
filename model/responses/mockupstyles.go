package responses

import model "github.com/baldurstod/go-printful-sdk/model"

type MockupStylesResponse struct {
	Data   []model.MockupStyles `json:"data" bson:"data"`
	Paging `json:"paging" bson:"paging"`
	Links  MockupStylesLinks `json:"_links" bson:"_links"`
}

type MockupStylesLinks struct {
	model.Links
	Product model.Link `json:"product" bson:"product"`
}
