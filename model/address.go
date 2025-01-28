package model

type Address struct {
	Name        string `json:"name" bson:"name"`
	Company     string `json:"company" bson:"company"`
	Address1    string `json:"address1" bson:"address1"`
	Address2    string `json:"address2" bson:"address2"`
	City        string `json:"city" bson:"city"`
	StateCode   string `json:"state_code" bson:"state_code"`
	StateName   string `json:"state_name" bson:"state_name"`
	CountryCode string `json:"country_code" bson:"country_code"`
	CountryName string `json:"country_name" bson:"country_name"`
	ZIP         string `json:"zip" bson:"zip"`
	Phone       string `json:"phone" bson:"phone"`
	Email       string `json:"email" bson:"email"`
	TaxNumber   string `json:"tax_number" bson:"tax_number"`
}
