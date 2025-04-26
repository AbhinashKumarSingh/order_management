package main

import (
	"example.com/app/apis"
	"example.com/app/cron_job"
	"github.com/labstack/echo/v4"
)

func addRoutes(e *echo.Echo) {
	// go order_service.Worker()
	v2 := e.Group("/sales-service")
	orderRoutes(v2)
	refreshRoutes(v2)
	cron_job.Init()
}

func orderRoutes(router *echo.Group) {
	routes := router.Group("/order")
	routes.POST("/revenue", apis.GetRevenueForDateRange)
	routes.POST("/revenue_by_products", apis.GetRevByProductForDateRange)
	routes.POST("/revenue_by_category", apis.GetRevByCategoryForDateRange)
	routes.POST("/revenue_by_region", apis.GetRevByRegionForDateRange)

}

func refreshRoutes(router *echo.Group) {
	routes := router.Group("/data")
	routes.GET("/refresh", apis.RefreshData)

}
