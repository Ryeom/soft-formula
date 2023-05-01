package main

import (
	"github.com/Ryeom/soft-formula/log"
	"github.com/Ryeom/soft-formula/router"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func init() {
	log.InitializeApplicationLog()
}
func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(log.GetCustomLogConfig()))
	router.Initialize(e)
	log.Logger.Fatal(e.Start(":8080"))

}
