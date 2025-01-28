package model

type CatalogItemType string

const (
	TypeOrderItem    CatalogItemType = "order_item"
	TypeBrandingItem CatalogItemType = "branding_item"
)

type CatalogItemSummary struct {
	ID               int                     `json:"id" bson:"id"`
	Type             CatalogItemType         `json:"Type" bson:"Type"`
	CatalogVariantID int                     `json:"catalog_variant_id" bson:"catalog_variant_id"`
	ExternalID       string                  `json:"external_id" bson:"external_id"`
	Quantity         int                     `json:"quantity" bson:"quantity"`
	Name             string                  `json:"name" bson:"name"`
	Price            string                  `json:"price" bson:"price"`
	RetailPrice      string                  `json:"retail_price" bson:"retail_price"`
	Currency         string                  `json:"currency" bson:"currency"`
	RetailCurrency   string                  `json:"retail_currency" bson:"retail_currency"`
	Links            CatalogItemSummaryLinks `json:"_links" bson:"_links"`
}

type CatalogItemSummaryLinks struct {
	Self Link `json:"self" bson:"self"`
}

func (o *CatalogItemSummary) isOrderItem() {}
