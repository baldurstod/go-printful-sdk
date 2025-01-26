package model

type File struct {
	ID           int       `json:"id" bson:"id"`
	URL          string    `json:"url" bson:"url"`
	Hash         string    `json:"hash" bson:"hash"`
	Filename     string    `json:"filename" bson:"filename"`
	MimeType     string    `json:"mime_type" bson:"mime_type"`
	Size         int       `json:"size" bson:"size"`
	Width        int       `json:"width" bson:"width"`
	Height       int       `json:"height" bson:"height"`
	Dpi          int       `json:"dpi" bson:"dpi"`
	Status       string    `json:"status" bson:"status"`
	Created      string    `json:"created" bson:"created"`
	ThumbnailURL string    `json:"thumbnail_url" bson:"thumbnail_url"`
	PreviewURL   string    `json:"preview_url" bson:"preview_url"`
	Visible      bool      `json:"visible" bson:"visible"`
	IsTemporary  bool      `json:"is_temporary" bson:"is_temporary"`
	Links        FileLinks `json:"_links" bson:"_links"`
}

type FileLinks struct {
	Self Link `json:"self" bson:"self"`
}
