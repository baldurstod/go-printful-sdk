package model

type PackingSlip struct {
	Email         string `json:"email" bson:"email"`
	Phone         string `json:"phone" bson:"phone"`
	Message       string `json:"message" bson:"message"`
	LogoURL       string `json:"logo_url" bson:"logo_url"`
	StoreName     string `json:"store_name" bson:"store_name"`
	CustomOrderID string `json:"custom_order_id" bson:"custom_order_id"`
}
