package stream

import (
	"stream/routes"

	"github.com/labstack/echo/v4"
)

func Init() {
	e := echo.New()

	// registers all the available routes
	routes.Init(e)

	// starts echo server
	e.Logger.Fatal(e.Start(":8082"))
}
