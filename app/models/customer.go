package models

type Customer struct {
	ID      int `gorm:"unique"`
	UUID    string
	Name    string
	Email   string
	Address string
}
