package responses

import model "github.com/baldurstod/go-printful-sdk/model"

type VariantImagesResponse struct {
	// Note: the schema suggest data is an array, but the reality differs
	Data model.VariantImages `json:"data" mapstructure:"data"`
}
