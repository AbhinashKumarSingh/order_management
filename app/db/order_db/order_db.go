package order_db

import (
	"example.com/app/config"
	"example.com/app/models"
	"example.com/app/response"
)

func GetRevenueForDateRange(startDate, endDate string) (*response.Revenue, error) {
	rev := new(response.Revenue)
	err := config.DB.Model(&models.OrderItem{}).
		Select("SUM(quantity * unit_price) as total_revenue").
		Joins("JOIN orders ON order_items.order_uuid = orders.uuid").
		Where("orders.date_of_sale BETWEEN ? AND ?", startDate, endDate).
		Scan(&rev).Error

	return rev, err
}

func GetTtlRevByProductForDateRange(startDate, endDate string) ([]*response.RevenueByProduct, error) {
	rev := make([]*response.RevenueByProduct, 0)

	err := config.DB.Table("order_items oi").
		Select("p.name as product_name, SUM((oi.unit_price * oi.quantity) - (oi.unit_price * oi.quantity * o.discount / 100)) AS total_revenue").
		Joins("JOIN products p ON oi.product_uuid = p.uuid").
		Joins("JOIN orders o ON oi.order_uuid = o.uuid").
		Where("o.date_of_sale BETWEEN ? AND ?", startDate, endDate).
		Group("p.uuid, p.name").
		Order("total_revenue DESC").
		Scan(&rev).Error

	return rev, err
}

func GetTtlRevByCategoryForDateRange(startDate, endDate string) ([]*response.RevenueByCategory, error) {
	rev := make([]*response.RevenueByCategory, 0)

	err := config.DB.Table("order_items oi").
		Select("p.category, SUM(oi.unit_price * oi.quantity) AS total_revenue").
		Joins("JOIN products p ON oi.product_uuid = p.uuid").
		Joins("JOIN orders o ON oi.order_uuid = o.uuid").
		Where("o.date_of_sale BETWEEN ? AND ?", startDate, endDate).
		Group("p.category").
		Order("total_revenue DESC").
		Scan(&rev).Error

	return rev, err
}

func GetTtlRevByRegionForDateRange(startDate, endDate string) ([]*response.RevenueByRegion, error) {
	rev := make([]*response.RevenueByRegion, 0)

	err := config.DB.Table("order_items oi").
		Select("r.region_name, SUM(oi.unit_price * oi.quantity) AS total_revenue").
		Joins("JOIN orders o ON oi.order_uuid = o.uuid").
		Joins("JOIN regions r ON o.uuid = r.order_uuid").
		Where("o.date_of_sale BETWEEN ? AND ?", startDate, endDate).
		Group("r.region_name").
		Order("total_revenue DESC").
		Scan(&rev).Error

	return rev, err
}
