package response

type Revenue struct {
	StartDate    string  `json:"start_date"`
	EndDate      string  `json:"end_date"`
	TotalRevenue float64 `json:"total_revenue"`
}
type RevenueByProduct struct {
	TotalRevenue float64 `json:"total_revenue"`
	ProductName  string  `json:"product_name"`
}

type RevenueByCategory struct {
	TotalRevenue float64 `json:"total_revenue"`
	Category     string  `json:"category"`
}

type RevenueByRegion struct {
	TotalRevenue float64 `json:"total_revenue"`
	RegionName   string  `json:"region"`
}
