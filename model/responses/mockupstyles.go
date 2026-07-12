package responses

import model "github.com/baldurstod/go-printful-sdk/model"

type MockupStylesResponse struct {
	Data   []model.MockupStyles `json:"data" bson:"data"`
	Paging `json:"paging" bson:"paging"`
}
