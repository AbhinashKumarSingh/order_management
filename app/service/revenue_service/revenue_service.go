package revenue_service

import (
	"fmt"

	"example.com/app/db/order_db"
	"example.com/app/response"
	"github.com/phuslu/log"
)

func GetRevenue(startDate, endDate string) (*response.Revenue, error) {

	rev, err := order_db.GetRevenueForDateRange(startDate, endDate)
	if err != nil {
		log.Error().Err(err).Msg("error in getting revenue")
		return nil, fmt.Errorf("Please try againg letter")
	}
	rev.StartDate = startDate
	rev.EndDate = endDate
	return rev, err
}

func GetRevByProuctForDateRange(startDate, endDate string) ([]*response.RevenueByProduct, error) {

	rev, err := order_db.GetTtlRevByProductForDateRange(startDate, endDate)
	if err != nil {
		log.Error().Err(err).Msg("error in getting revenue")
		return nil, fmt.Errorf("Please try againg letter")
	}
	return rev, err
}

func GetRevByCategoryForDateRange(startDate, endDate string) ([]*response.RevenueByCategory, error) {

	rev, err := order_db.GetTtlRevByCategoryForDateRange(startDate, endDate)
	if err != nil {
		log.Error().Err(err).Msg("error in getting revenue")
		return nil, fmt.Errorf("Please try againg letter")
	}
	return rev, err
}

func GetRevByRegionForDateRange(startDate, endDate string) ([]*response.RevenueByRegion, error) {

	rev, err := order_db.GetTtlRevByRegionForDateRange(startDate, endDate)
	if err != nil {
		log.Error().Err(err).Msg("error in getting revenue")
		return nil, fmt.Errorf("Please try againg letter")
	}
	return rev, err
}
