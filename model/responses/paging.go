package responses

type Paging struct {
	Total  uint `json:"total" bson:"total" mapstructure:"total"`
	Offset uint `json:"offset" bson:"offset" mapstructure:"offset"`
	Limit  uint `json:"limit" bson:"limit" mapstructure:"limit"`
}
