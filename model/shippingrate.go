package model

type ShippingRate struct {
	Shipping           string      `json:"shipping" bson:"shipping"`
	ShippingMethodName string      `json:"shipping_method_name" bson:"shipping_method_name"`
	Rate               string      `json:"rate" bson:"rate"`
	Currency           string      `json:"currency" bson:"currency"`
	MinDeliveryDays    int         `json:"min_delivery_days" bson:"min_delivery_days"`
	MaxDeliveryDays    int         `json:"max_delivery_days" bson:"max_delivery_days"`
	MinDeliveryDate    string      `json:"min_delivery_date" bson:"min_delivery_date"`
	MaxDeliveryDate    string      `json:"max_delivery_date" bson:"max_delivery_date"`
	Shipments          []Shipment2 `json:"shipments" bson:"shipments"`
}

type Shipment2 struct {
	DepartureCountry    string          `json:"departure_country" bson:"departure_country"`
	ShipmentItems       []ShipmentItem2 `json:"shipment_items" bson:"shipment_items"`
	CustomsFeesPossible bool            `json:"customs_fees_possible" bson:"customs_fees_possible"`
}

type ShipmentItem2 struct {
	CatalogVariantID int `json:"catalog_variant_id" bson:"catalog_variant_id"`
	Quantity         int `json:"quantity" bson:"quantity"`
}
