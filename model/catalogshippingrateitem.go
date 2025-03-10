package model

type CatalogShippingRateItem struct {
	Source           string `json:"source" bson:"source"`
	Quantity         int    `json:"quantity" bson:"quantity"`
	CatalogVariantID int    `json:"catalog_variant_id" bson:"catalog_variant_id"`
}

func (c CatalogShippingRateItem) IsCatalogOrWarehouseShippingRateItem() {}
