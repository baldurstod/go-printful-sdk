package model

type Links struct {
	Self     Link `json:"self" bson:"self" mapstructure:"self"`
	Next     Link `json:"next" bson:"next" mapstructure:"next"`
	Previous Link `json:"previous" bson:"previous" mapstructure:"previous"`
	First    Link `json:"first" bson:"first" mapstructure:"first"`
	Last     Link `json:"last" bson:"last" mapstructure:"last"`
}

type Link struct {
	Href string `json:"href" bson:"href" mapstructure:"href"`
}
