package responses

import model "github.com/baldurstod/go-printful-sdk/model"

type VariantPricesResponse struct {
	Data model.VariantPrice `json:"data" bson:"data" mapstructure:"data"`
}
