package model

type Address struct {
	Name        string `json:"name" bson:"name"  mapstructure:"name"`
	Company     string `json:"company" bson:"company"  mapstructure:"company"`
	Address1    string `json:"address1" bson:"address1"  mapstructure:"address1"`
	Address2    string `json:"address2" bson:"address2"  mapstructure:"address2"`
	City        string `json:"city" bson:"city"  mapstructure:"city"`
	StateCode   string `json:"state_code" bson:"state_code"  mapstructure:"state_code"`
	StateName   string `json:"state_name" bson:"state_name"  mapstructure:"state_name"`
	CountryCode string `json:"country_code" bson:"country_code"  mapstructure:"country_code"`
	CountryName string `json:"country_name" bson:"country_name"  mapstructure:"country_name"`
	ZIP         string `json:"zip" bson:"zip"  mapstructure:"zip"`
	Phone       string `json:"phone" bson:"phone"  mapstructure:"phone"`
	Email       string `json:"email" bson:"email"  mapstructure:"email"`
	TaxNumber   string `json:"tax_number" bson:"tax_number"  mapstructure:"tax_number"`
}
