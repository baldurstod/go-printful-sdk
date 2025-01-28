package model

type Item struct {
	ID             int            `json:"id" bson:"id"`
	ExternalID     string         `json:"external_id,omitempty" bson:"external_id"`
	Quantity       int            `json:"quantity" bson:"quantity"`
	RetailPrice    string         `json:"retail_price" bson:"retail_price"`
	Name           string         `json:"name" bson:"name"`
	Placements     PlacementsList `json:"placements" bson:"placements"`
	ProductOptions `json:"product_options" bson:"product_options"`
}

type PlacementsList = []Placement
