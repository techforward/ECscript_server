package routes

import (
	"github.com/techforward/ECscript_server/handler"

	"github.com/labstack/echo"
)

// OtherRouter OtherRouter
func OtherRouter(e *echo.Echo) {
	g := e.Group("/other")

	g.GET("", handler.GetAllOthers)
	g.GET(":id", handler.GetOther)

	g.POST("", handler.CreateOther)
	g.PUT(":id", handler.UpdateOther)
	g.DELETE(":id", handler.DeleteOther)

	// g.POST("", handler.CreateOther, middleware.JWT([]byte("secret")))
	// g.PUT(":id", handler.UpdateOther, middleware.JWT([]byte("secret")))
	// g.DELETE(":id", handler.DeleteOther, middleware.JWT([]byte("secret")))
}
