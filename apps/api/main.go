package main

import (
	"fmt"
	"go-next-template/api/env"
	"go-next-template/api/router"

	"github.com/labstack/echo/v4/middleware"
)

func main() {
	api_env.ParseApiEnv()

	// ApiRoot with Echo instance
	r := api_router.CreateApiRouter()

	r.Echo().Use(middleware.CORS())

	// Start server
	r.Echo().Logger.Fatal(r.Echo().Start(fmt.Sprintf(":%d", api_env.API_PORT)))
}
