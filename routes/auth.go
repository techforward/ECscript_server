package routes

import (
	"github.com/techforward/ECscript_server/handler"

	"github.com/labstack/echo"
)

// AuthRouter AuthRouter
func AuthRouter(e *echo.Echo) {
	g := e.Group("/auth")

	g.GET("/:idToken", handler.VeryfyIDToken)
}
