package model

type CatalogItem struct {
	Source           string `json:"source" bson:"source"`
	CatalogVariantID int    `json:"catalog_variant_id" bson:"catalog_variant_id"`
	Item             `json:"item" bson:"item"`
}
