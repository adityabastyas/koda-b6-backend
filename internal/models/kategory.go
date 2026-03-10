package models

type Kategory struct {
	KategoryID int    `json:"kategory_id" db:"kategory_id"`
	Name       string `json:"name" db:"name"`
}

type KategoryInput struct {
	Name string `json:"name"`
}
