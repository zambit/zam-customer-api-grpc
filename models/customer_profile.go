package models

import "time"

type Profile struct {
	tableName struct{}  `sql:"personal_data"`
	ID        uint64    `json:"id"`
	UserID    uint64    `json:"user_id"`
	StatusID  uint64    `json:"status_id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	BirthDate time.Time `json:"birth_date"`
	Gender    string    `json:"gender"`
	Country   string    `json:"country"`
	Address   string    `json:"address"`
}
