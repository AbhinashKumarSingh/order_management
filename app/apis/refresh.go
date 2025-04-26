package apis

import (
	"net/http"

	"example.com/app/constants"
	"example.com/app/utils"
	"github.com/labstack/echo/v4"
)

func RefreshData(c echo.Context) error {
	err := utils.LoadCSV(constants.CSVFilePath)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})

	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Data refreshed successfully"})
}
