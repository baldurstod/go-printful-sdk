package responses

import model "github.com/baldurstod/go-printful-sdk/model"

type CreateOrderResponse struct {
	Data model.Order `json:"data" bson:"data"`
}
