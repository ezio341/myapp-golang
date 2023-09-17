package routes

import (
	"myproject/controllers"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoutesDepartment(e *echo.Echo) {
	auth := e.Group("")
	auth.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_KEY"))))
	auth.GET("/departments", controllers.GetDepartment)
	auth.POST("/department", controllers.AddDepartment)
	auth.DELETE("/department/:id", controllers.DeleteDepartment)
}
