package model

type Category struct {
	ID       int    `json:"id" bson:"id" mapstructure:"id"`
	ParentID int    `json:"parent_id" bson:"parent_id" mapstructure:"parent_id"`
	ImageURL string `json:"image_url" bson:"image_url" mapstructure:"image_url"`
	Title    string `json:"title" bson:"title" mapstructure:"title"`
}
