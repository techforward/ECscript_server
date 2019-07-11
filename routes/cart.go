package routes

import (
	"github.com/techforward/ECscript_server/handler"

	"github.com/labstack/echo"
)

// CartRouter CartRouter
func CartRouter(e *echo.Echo) {
	g := e.Group("/cart")

	g.GET("/", handler.GetAllCarts)
	g.GET("/:id", handler.GetCart)

	g.POST("/", handler.CreateCart)
	g.PUT("/:id", handler.UpdateCart)
	g.DELETE("/:id", handler.DeleteCart)

	// loginRequired := e.Group("/cart")
	// loginRequired.Use(middleware.JWT([]byte("secret")))
	// loginRequired.POST("/", handler.CreateCart)
	// loginRequired.PUT("/:id", handler.UpdateCart)
	// loginRequired.DELETE("/:id", handler.DeleteCart)
}
