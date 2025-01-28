package model

type CalculationStatus string

const (
	Done        CalculationStatus = "done"
	Calculating CalculationStatus = "calculating"
	Failed      CalculationStatus = "failed"
)

type Costs struct {
	CalculationStatus `json:"calculation_status" bson:"calculation_status"`
	Currency          string `json:"currency" bson:"currency"`
	Subtotal          string `json:"subtotal" bson:"subtotal"`
	Discount          string `json:"discount" bson:"discount"`
	Shipping          string `json:"shipping" bson:"shipping"`
	Digitization      string `json:"digitization" bson:"digitization"`
	AdditionalFee     string `json:"additional_fee" bson:"additional_fee"`
	FulfillmentFee    string `json:"fulfillment_fee" bson:"fulfillment_fee"`
	RetailDeliveryFee string `json:"retail_delivery_fee" bson:"retail_delivery_fee"`
	Vat               string `json:"vat" bson:"vat"`
	Tax               string `json:"tax" bson:"tax"`
	Total             string `json:"total" bson:"total"`
}
