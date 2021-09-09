package entity

type Response struct {
	City string 			`bson:"city" json:"city"`
	Value string			`bson:"value" json:"value"`
	TimeRequested string	`bson:"time_requested" json:"time_requested"`
}
