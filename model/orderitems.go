package model

type OrderItems []OrderItem

type OrderItemItemSource string

const (
	SourceCatalog   OrderItemItemSource = "catalog"
	SourceWarehouse OrderItemItemSource = "warehouse"
)

type OrderItem struct {
	Source string `json:"source" bson:"source"`
	CatalogItemSummary
	//TODO: add WarehouseItemSummary
}
