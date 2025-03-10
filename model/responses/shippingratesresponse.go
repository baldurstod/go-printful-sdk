package responses

import model "github.com/baldurstod/go-printful-sdk/model"

type ShippingRatesResponse struct {
	Data []model.ShippingRate `json:"data" bson:"data"`
}
