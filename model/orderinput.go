package model

type OrderInput struct {
	ExternalID    string        `json:"external_id" bson:"external_id" mapstructure:"external_id"`
	Shipping      string        `json:"shipping" bson:"shipping" mapstructure:"shipping"`
	Recipient     Address       `json:"recipient" bson:"recipient" mapstructure:"recipient"`
	OrderItems    []CatalogItem `json:"order_items" bson:"order_items" mapstructure:"order_items"`
	Customization Customization `json:"customization" bson:"customization" mapstructure:"customization"`
	RetailCosts   RetailCosts2  `json:"retail_costs" bson:"retail_costs" mapstructure:"retail_costs"`
}
