package model

type Customization struct {
	Gift        `json:"gift" bson:"gift"`
	PackingSlip `json:"packing_slip" bson:"packing_slip"`
}
