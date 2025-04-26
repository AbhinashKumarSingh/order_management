package apis

import (
	"net/http"

	"example.com/app/service/revenue_service"
	"github.com/labstack/echo/v4"
	"github.com/phuslu/log"
)

func GetRevenueForDateRange(c echo.Context) error {
	startDate := c.QueryParam("start_date")
	endDate := c.QueryParam("end_date")

	if startDate == "" {
		log.Error().Msg("GetRevenue: startDate is not valid")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}
	if endDate == "" {
		log.Error().Msg("GetRevenue: endDate is not valid")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	rev, err := revenue_service.GetRevenue(startDate, endDate)
	if err != nil {
		log.Error().Err(err).Msg("GetOrderStatus: Failed to get order")
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get result"})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"revenue": rev})
}

func GetRevByProductForDateRange(c echo.Context) error {
	startDate := c.QueryParam("start_date")
	endDate := c.QueryParam("end_date")

	if startDate == "" {
		log.Error().Msg("GetRevenue: startDate is not valid")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}
	if endDate == "" {
		log.Error().Msg("GetRevenue: endDate is not valid")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	rev, err := revenue_service.GetRevByProuctForDateRange(startDate, endDate)
	if err != nil {
		log.Error().Err(err).Msg("GetOrderStatus: Failed to get order")
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get result"})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"revenue": rev})
}

func GetRevByCategoryForDateRange(c echo.Context) error {
	startDate := c.QueryParam("start_date")
	endDate := c.QueryParam("end_date")

	if startDate == "" {
		log.Error().Msg("GetRevenue: startDate is not valid")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}
	if endDate == "" {
		log.Error().Msg("GetRevenue: endDate is not valid")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	rev, err := revenue_service.GetRevByCategoryForDateRange(startDate, endDate)
	if err != nil {
		log.Error().Err(err).Msg("GetOrderStatus: Failed to get order")
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get result"})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"revenue": rev})
}

func GetRevByRegionForDateRange(c echo.Context) error {
	startDate := c.QueryParam("start_date")
	endDate := c.QueryParam("end_date")

	if startDate == "" {
		log.Error().Msg("GetRevenue: startDate is not valid")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}
	if endDate == "" {
		log.Error().Msg("GetRevenue: endDate is not valid")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	rev, err := revenue_service.GetRevByRegionForDateRange(startDate, endDate)
	if err != nil {
		log.Error().Err(err).Msg("GetOrderStatus: Failed to get order")
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get result"})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"revenue": rev})
}
