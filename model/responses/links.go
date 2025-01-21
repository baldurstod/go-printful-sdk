package responses

import model "github.com/baldurstod/go-printful-sdk/model"

type Links struct {
	Self     model.Link `json:"self" bson:"self" mapstructure:"self"`
	Next     model.Link `json:"next" bson:"next" mapstructure:"next"`
	Previous model.Link `json:"previous" bson:"previous" mapstructure:"previous"`
	First    model.Link `json:"first" bson:"first" mapstructure:"first"`
	Last     model.Link `json:"last" bson:"last" mapstructure:"last"`
}
