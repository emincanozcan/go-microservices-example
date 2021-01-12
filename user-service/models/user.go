package models

type User struct {
	ID       uint   `json:"id" gorm:"primarykey"`
	Type     uint   `json:"type" gorm:"default:0"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}
