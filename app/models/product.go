package models

type Product struct {
	ID          int `gorm:"unique"`
	UUID        string
	Name        string
	Category    string
	Description string
}
