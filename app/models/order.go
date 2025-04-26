package models

type Order struct {
	ID            int `gorm:"unique"`
	UUID          string
	CustomerUUID  string
	DateOfSale    string
	PaymentMethod string
	ShippingCost  float64
	Discount      float64
}
