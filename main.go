package main

import (
	"context"
	"errors"
	"log"
	"myproject/configs"
	"myproject/routes"
	"net/http"
	"os"
	"os/signal"

	"github.com/labstack/echo/v4"
)

func ErrorMap() map[error]int {
	errorCodeMaps := make(map[error]int)
	errorCodeMaps[errors.New("Unauthorized")] = http.StatusUnauthorized
	return errorCodeMaps
}

func main() {
	var srv http.Server
	// close idle connection
	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		// received an interrupt signal, shut down.
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()
	e := echo.New()
	// configs.LoadEnv()
	configs.InitDatabase()
	routes.InitRoutesUser(e)
	routes.InitRoutesRole(e)
	routes.InitRoutesPosition(e)
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
