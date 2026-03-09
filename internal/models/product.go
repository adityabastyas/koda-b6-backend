package models

type Product struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       int      `json:"price"`
	Category    string   `json:"category"`
	Image       []string `json:"image"`
}
