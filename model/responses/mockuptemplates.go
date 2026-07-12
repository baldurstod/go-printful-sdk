package responses

import model "github.com/baldurstod/go-printful-sdk/model"

type MockupTemplatesResponse struct {
	Data   []model.MockupTemplates `json:"data" bson:"data"`
	Paging `json:"paging" bson:"paging"`
}
