package routes

import (
	"myproject/controllers"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoutesUser(e *echo.Echo) {
	e.POST("/login", controllers.Login)

	auth := e.Group("")
	auth.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_KEY"))))
	auth.POST("/user", controllers.AddUser)
	auth.PUT("/user/:id", controllers.EditUser)
	auth.DELETE("/user/:id", controllers.DeleteUser)
	auth.GET("/users", controllers.GetUsers)
	auth.GET("/user/:id", controllers.GetUser)
	auth.POST("/role", controllers.AddRole)

}
