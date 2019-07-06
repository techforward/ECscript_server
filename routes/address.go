package routes

import (
	"github.com/techforward/ECscript_server/handler"

	"github.com/labstack/echo"
)

func AddressRouter(e *echo.Echo) {
	g := e.Group("/address")

	g.GET("", handler.GetAllAddresses)
	g.GET(":id", handler.GetAddress)

	g.POST("", handler.CreateAddress)
	g.PUT(":id", handler.UpdateAddress)
	g.DELETE(":id", handler.DeleteAddress)

	// g.POST("", handler.CreateAddress, middleware.JWT([]byte("secret")))
	// g.PUT(":id", handler.UpdateAddress, middleware.JWT([]byte("secret")))
	// g.DELETE(":id", handler.DeleteAddress, middleware.JWT([]byte("secret")))
}
