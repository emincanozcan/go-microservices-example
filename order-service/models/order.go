package models

import "time"

type Order struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	UserID    uint      `json:"user_id"`
	Price     float32   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	Items     []Item    `json:"items"`
}
type Item struct {
	ID        uint    `json:"-" gorm:"primarykey"`
	OrderID   uint    `json:"order_id"`
	ProductID uint    `json:"product_id"`
	Title     string  `json:"title"`
	Count     uint    `json:"count"`
	Price     float32 `json:"price"`
}
