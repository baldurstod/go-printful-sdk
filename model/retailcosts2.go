package model

type RetailCosts2 struct {
	Currency string `json:"currency" bson:"currency"`
	Discount string `json:"discount" bson:"discount"`
	Shipping string `json:"shipping" bson:"shipping"`
	Tax      string `json:"tax" bson:"tax"`
}
