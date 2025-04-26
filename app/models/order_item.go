package models

type OrderItem struct {
	ID          int
	UUID        string
	OrderUUID   string
	ProductUUID string
	Quantity    int
	UnitPrice   float64

	Product *Product `gorm:"foreignKey:ProductUUID;references:UUID"`
	Order   *Order   `gorm:"foreignKey:OrderUUID;references:UUID"`
}
