package responses

import model "github.com/baldurstod/go-printful-sdk/model"

type MockupTemplatesResponse struct {
	Data   []model.MockupTemplates `json:"data" bson:"data"`
	Paging `json:"paging" bson:"paging"`
	Links  MockupTemplatesLinks `json:"_links" bson:"_links"`
}

type MockupTemplatesLinks struct {
	model.Links
	Product model.Link `json:"product" bson:"product"`
}
