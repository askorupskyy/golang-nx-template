package api_router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pangpanglabs/echoswagger/v2"
)

type User struct {
	Name string `json:"name"`
}

func createUser(c echo.Context) error {
	return c.JSON(http.StatusCreated, nil)
}

func getUserById(c echo.Context) error {
	return c.JSON(http.StatusOK, User{
		Name: "Test User",
	})
}

func getUsers(c echo.Context) error {
	users := []User{
		{
			Name: "test user 1",
		},
		{
			Name: "test user 2",
		},
	}

	return c.JSON(http.StatusOK, users)
}

func CreateApiRouter() echoswagger.ApiRoot {
	e := echo.New()
	r := echoswagger.New(e, "/doc", nil)

	g := r.Group("Users", "/users")

	g.GET("/", getUsers).
		SetOperationId("getUsers").
		AddResponse(http.StatusOK, "successful", []User{}, nil)

	g.GET("/:id", getUserById).
		SetOperationId("getUserById").
		AddParamPath("", "id", "User Id").
		AddResponse(http.StatusOK, "successful", User{}, nil)

	g.POST("/", createUser).
		SetOperationId("createUser").
		AddParamBody(User{}, "body", "User input struct", true).
		AddResponse(http.StatusCreated, "successful", nil, nil)

	return r
}
