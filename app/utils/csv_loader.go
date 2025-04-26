package utils

import (
	"encoding/csv"

	"os"
	"strconv"

	"example.com/app/config"
	"example.com/app/models"
	"github.com/google/uuid"
	"github.com/phuslu/log"
)

func LoadCSV(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	for idx, row := range records {
		if idx == 0 {
			continue // skip header
		}

		customer := models.Customer{
			UUID:    row[2],
			Name:    row[12],
			Email:   row[13],
			Address: row[14],
		}
		config.DB.FirstOrCreate(&customer, models.Customer{ID: customer.ID})

		product := models.Product{
			UUID:     row[1],
			Name:     row[3],
			Category: row[4],
		}
		config.DB.FirstOrCreate(&product, models.Product{ID: product.ID})

		order := models.Order{
			ID:            parseInt(row[0]),
			UUID:          uuid.New().String(),
			CustomerUUID:  customer.UUID,
			DateOfSale:    row[6],
			PaymentMethod: row[11],
			ShippingCost:  parseFloat(row[10]),
			Discount:      parseFloat(row[9]),
		}
		config.DB.FirstOrCreate(&order, models.Order{ID: order.ID})

		orderItem := models.OrderItem{
			UUID:        uuid.New().String(),
			OrderUUID:   order.UUID,
			ProductUUID: product.UUID,
			Quantity:    parseInt(row[7]),
			UnitPrice:   parseFloat(row[8]),
		}
		config.DB.Create(&orderItem)

		region := models.Region{
			OrderUUID:  order.UUID,
			RegionName: row[5],
		}
		config.DB.FirstOrCreate(&region, models.Region{OrderUUID: region.OrderUUID})
	}

	log.Info().Msg("csv file loadded successfully")

	return nil
}

func parseFloat(s string) float64 {
	val, _ := strconv.ParseFloat(s, 64)
	return val
}

func parseInt(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}
