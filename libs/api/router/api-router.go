package api_router

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pangpanglabs/echoswagger/v2"
)

type User struct {
	Name string `json:"name"`
}

// Handler
func createUser(c echo.Context) error {
	log.Println("creating user...")
	return c.JSON(http.StatusCreated, nil)
}

func CreateApiRouter() echoswagger.ApiRoot {
	e := echo.New()
	r := echoswagger.New(e, "/doc", nil)

	// Routes with parameters & responses
	r.POST("/users", createUser).
		SetOperationId("createUser").
		AddParamBody(User{}, "body", "User input struct", true).
		AddResponse(http.StatusCreated, "successful", nil, nil)

	return r
}
