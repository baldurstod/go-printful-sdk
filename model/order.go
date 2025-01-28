package model

type Order struct {
	ID          int     `json:"id" bson:"id"`
	ExternalID  string  `json:"external_id" bson:"external_id"`
	StoreID     int     `json:"store_id" bson:"store_id"`
	Shipping    string  `json:"shipping" bson:"shipping"`
	Status      string  `json:"status" bson:"status"`
	CreatedAt   string  `json:"created_at" bson:"created_at"`
	UpdatedAt   string  `json:"updated_at" bson:"updated_at"`
	Recipient   Address `json:"recipient" bson:"recipient"`
	Costs       `json:"costs" bson:"costs"`
	RetailCosts `json:"retail_costs" bson:"retail_costs"`
	OrderItems  `json:"order_items" bson:"order_items"`
}
