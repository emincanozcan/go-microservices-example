package models

type Product struct {
	ID          uint    `json:"id" gorm:"primary_key"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Stock       uint    `json:"stock"`
	Price       float32 `json:"price"`
}
