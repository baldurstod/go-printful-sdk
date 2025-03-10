package model

type ShippingRatesAddress struct {
	Address1    string `json:"address1" bson:"address1" mapstructure:"address1"`
	Address2    string `json:"address2" bson:"address2" mapstructure:"address2"`
	City        string `json:"city" bson:"city" mapstructure:"city"`
	StateCode   string `json:"state_code" bson:"state_code" mapstructure:"state_code"`
	CountryCode string `json:"country_code" bson:"country_code" mapstructure:"country_code"`
	ZIP         string `json:"zip" bson:"zip" mapstructure:"zip"`
}
