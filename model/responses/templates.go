package responses

import model "github.com/baldurstod/go-printful-sdk/model"

// TODO: adapt for api v2
type TemplatesResponse struct {
	Code   int                    `json:"code" bson:"code" mapstructure:"code"`
	Result model.ProductTemplates `json:"result" bson:"result" mapstructure:"result"`
}
