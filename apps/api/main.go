package main

import (
	"github.com/askorupskyy/go-next-template/libs/api/router"

	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// ApiRoot with Echo instance
	r := api_router.CreateApiRouter()

	r.Echo().Use(middleware.CORS())

	// Start server
	r.Echo().Logger.Fatal(r.Echo().Start(":1323"))
}
