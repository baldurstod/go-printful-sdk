package model

type CatalogOrWarehouseShippingRateItem struct {
	Source             string `json:"source" bson:"source" mapstructure:"source"`
	Quantity           int    `json:"quantity" bson:"quantity" mapstructure:"quantity"`
	CatalogVariantID   int    `json:"catalog_variant_id" bson:"catalog_variant_id" mapstructure:"catalog_variant_id"`
	WarehouseVariantID int    `json:"warehouse_variant_id" bson:"warehouse_variant_id" mapstructure:"warehouse_variant_id"`
}
