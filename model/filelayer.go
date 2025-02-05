package model

type FileLayer struct {
	Type         string          `json:"type" bson:"type"`
	LayerOptions []CatalogOption `json:"layer_options" bson:"layer_options"`
}
