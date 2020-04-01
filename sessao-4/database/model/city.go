package model

import "database/sql"

type City struct {
	Country string `json:"country" db:"country"`
	City sql.NullString `json:"city" db:"city"`
	PhoneCode int `json:"phone_code" db:"phone_code"`
}
