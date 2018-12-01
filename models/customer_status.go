package models

type CustomerStatus struct {
	tableName struct{} `sql:"customer_statuses"`
	ID        uint64   `json:"id"`
	Name      string   `json:"name"`
}
