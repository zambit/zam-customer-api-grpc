package models

import (
	"time"
)

type Customer struct {
	tableName    struct{}  `sql:"customers"`
	ID           uint64    `json:"id"`
	Phone        string    `json:"phone"`
	Password     string    `json:"-"`
	RegisteredAt time.Time `json:"registered_at"`
	ReferrerID   uint64    `json:"referrer_id"`
	StatusID     uint64    `json:"status_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
