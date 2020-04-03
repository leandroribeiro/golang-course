package model

type City struct {
	Country string `json:"country" db:"country" bson:"country"`
	City string `json:"city" db:"city" bson:"city"`
	PhoneCode int `json:"phone_code" db:"phone_code" bson:"phone_code"`
}
