package model

type Gift struct {
	Subject string `json:"subject" bson:"subject"`
	Message string `json:"message" bson:"message"`
}
