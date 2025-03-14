package responses

import model "github.com/baldurstod/go-printful-sdk/model"

type CreateOrderResponse struct {
	Data model.Order `json:"data" bson:"data"`
}

type GetItemById struct {
	Data model.CatalogItemReadonly `json:"data" bson:"data" mapstructure:"data"`
}
