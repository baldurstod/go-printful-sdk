package responses

import model "github.com/baldurstod/go-printful-sdk/model"

type AddFileResponse struct {
	Data model.File `json:"data" bson:"data"`
}
