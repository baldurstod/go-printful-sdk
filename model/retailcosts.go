package model

type RetailCosts struct {
	CalculationStatus `json:"calculation_status" bson:"calculation_status"`
	Currency          string `json:"currency" bson:"currency"`
	Subtotal          string `json:"subtotal" bson:"subtotal"`
	Discount          string `json:"discount" bson:"discount"`
	Shipping          string `json:"shipping" bson:"shipping"`
	Vat               string `json:"vat" bson:"vat"`
	Tax               string `json:"tax" bson:"tax"`
	Total             string `json:"total" bson:"total"`
}
