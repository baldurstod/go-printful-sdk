package model

type Category struct {
	ID       int           `json:"id" bson:"id" mapstructure:"id"`
	ParentID int           `json:"parent_id" bson:"parent_id" mapstructure:"parent_id"`
	ImageURL string        `json:"image_url" bson:"image_url" mapstructure:"image_url"`
	Title    string        `json:"title" bson:"title" mapstructure:"title"`
	Links    CategoryLinks `json:"_links" bson:"_links" mapstructure:"_links"`
}

type CategoryLinks struct {
	Self Link `json:"self" bson:"self" mapstructure:"self"`
}
