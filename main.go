package main

import (
	"errors"
	"myproject/configs"
	"myproject/routes"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func ErrorMap() map[error]int {
	errorCodeMaps := make(map[error]int)
	errorCodeMaps[errors.New("Unauthorized")] = http.StatusUnauthorized
	return errorCodeMaps
}

func main() {
	e := echo.New()
	configs.LoadEnv()
	configs.InitDatabase()
	routes.InitRoutesUser(e)
	routes.InitRoutesRole(e)
	routes.InitRoutesDepartment(e)
	e.Start(GetPort())
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "3001"
	}
	return ":" + port
}
