package routes

import (
	"github.com/techforward/ECscript_server/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func UserRouter(e *echo.Echo) {
	g := e.Group("/user")

	g.GET("/", handler.GetAllUsers)
	g.GET("/:id", handler.GetUser)
	g.POST("/", handler.CreateUser)

	loginRequired := e.Group("/user")
	loginRequired.Use(middleware.JWT([]byte("secret")))
	loginRequired.PUT("/:id", handler.UpdateUser)
	loginRequired.DELETE("/:id", handler.DeleteUser)
}
