package routes

import (
	"myproject/controllers"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoutesPosition(e *echo.Echo) {
	auth := e.Group("")
	auth.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_KEY"))))
	auth.GET("/positions", controllers.GetPositions)
	auth.POST("/position", controllers.AddPosition)
	auth.DELETE("/position/:id", controllers.DeletePosition)
}
