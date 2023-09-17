package routes

import (
	"myproject/controllers"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoutesRole(e *echo.Echo) {
	auth := e.Group("")
	auth.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_KEY"))))
	auth.GET("/roles", controllers.GetRoles)
	auth.POST("/role", controllers.AddRole)
	auth.DELETE("/role/:id", controllers.DeleteRole)

}
