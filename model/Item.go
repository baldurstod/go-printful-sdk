package model

type Item struct {
	ID             int            `json:"id" bson:"id" mapstructure:"id"`
	ExternalID     string         `json:"external_id,omitempty" bson:"external_id" mapstructure:"external_id"`
	Quantity       int            `json:"quantity" bson:"quantity" mapstructure:"quantity"`
	RetailPrice    string         `json:"retail_price" bson:"retail_price" mapstructure:"retail_price"`
	Name           string         `json:"name" bson:"name" mapstructure:"name"`
	Placements     PlacementsList `json:"placements" bson:"placements" mapstructure:"placements"`
	ProductOptions `json:"product_options,omitempty" bson:"product_options" mapstructure:"product_options"`
}

type PlacementsList = []Placement

type ItemReadonly = Item
