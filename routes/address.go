package routes

import (
	"github.com/techforward/ECscript_server/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func AddressRouter(e *echo.Echo) {
	g := e.Group("/address")

	g.GET("/", handler.GetAllAddresses)
	g.GET("/:id", handler.GetAddress)

	loginRequired := e.Group("/address")

	loginRequired.Use(middleware.JWT([]byte("secret")))
	loginRequired.POST("/", handler.CreateAddress)
	loginRequired.PUT("/:id", handler.UpdateAddress)
	loginRequired.DELETE("/:id", handler.DeleteAddress)
}
