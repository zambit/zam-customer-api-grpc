package models

type ProfileStatus struct {
	tableName struct{} `sql:"profile_data_statuses"`
	ID        uint64   `json:"id"`
	Name      string   `json:"name"`
}
