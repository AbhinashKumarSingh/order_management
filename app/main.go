package main

import (
	"example.com/app/config"
	"github.com/labstack/echo/v4"
)

var Echo *echo.Echo

func main() {
	config.InitDB()
	Echo := echo.New()
	addRoutes(Echo)
	Echo.Logger.Fatal(Echo.Start(":8000"))
}
