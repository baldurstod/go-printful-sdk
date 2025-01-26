package responses

type Error4XXResponse struct {
	Code   int    `json:"code" bson:"code"`
	Result string `json:"result" bson:"result"`
	Error  `json:"error" bson:"error"`
}

type Error struct {
	Reason  string `json:"reason" bson:"reason"`
	Message string `json:"message" bson:"message"`
}

type Error5XXResponse struct {
	Type     string `json:"type" bson:"type"`
	Status   int    `json:"status" bson:"status"`
	Title    string `json:"title" bson:"title"`
	Details  string `json:"details" bson:"details"`
	Instance string `json:"instance" bson:"instance"`
}
