package api_env

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Env struct {
	API_ROOT_URL string
	APP_ROOT_URL string
}

var API_PORT int
var API_ROOT_URL string
var APP_ROOT_URL string

func ParseApiEnv() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	env, err := godotenv.Read(wd + "/../../.env")
	if err != nil {
		log.Fatal(err)
	}

	parsed := Env{
		API_ROOT_URL: env["API_ROOT_URL"],
		APP_ROOT_URL: env["APP_ROOT_URL"],
	}

	if parsed.APP_ROOT_URL == "" || parsed.API_ROOT_URL == "" {
		log.Fatal("Failed to parse .env")
	}

	split_url := strings.Split(parsed.API_ROOT_URL, ":")
	if len(split_url) <= 1 {
		API_PORT = 3000
	}

	parsed_port, err := strconv.Atoi(split_url[len(split_url)-1])
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to parse the port from URL: %s", split_url))
	}

	API_PORT = parsed_port
	API_ROOT_URL = parsed.API_ROOT_URL
	APP_ROOT_URL = parsed.APP_ROOT_URL
}
