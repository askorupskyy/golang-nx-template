package main

import (
	"encoding/json"
	"log"
	"net/http/httptest"
	"os"

	"github.com/askorupskyy/go-next-template/libs/api/router"
	"github.com/labstack/echo"
	"github.com/pangpanglabs/echoswagger/v2"
)

func main() {
	r := api_router.CreateApiRouter()
	f, err := os.Create("./schema.json")

	req := httptest.NewRequest(echo.GET, "/doc", nil)
	rec := httptest.NewRecorder()

	spec, err := r.(*echoswagger.Root).GetSpec(r.Echo().NewContext(req, rec), "/doc")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	json, err := json.MarshalIndent(spec, "", "\t")

	if err != nil {
		log.Fatal(err)
	}

	_, err = f.Write(json)

	os.Exit(0)
}
