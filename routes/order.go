package routes

import (
	"github.com/techforward/ECscript_server/handler"

	"github.com/labstack/echo"
)

// OrderRouter OrderRouter
func OrderRouter(e *echo.Echo) {
	g := e.Group("/order")

	g.GET("/", handler.GetAllOrders)
	g.GET("/:id", handler.GetOrder)

	g.POST("/", handler.CreateOrder)
	g.PUT("/:id", handler.UpdateOrder)
	g.DELETE("/:id", handler.DeleteOrder)

	// loginRequired := e.Group("/order")
	// loginRequired.Use(middleware.JWT([]byte("secret")))
	// loginRequired.POST("/", handler.CreateOrder)
	// loginRequired.PUT("/:id", handler.UpdateOrder)
	// loginRequired.DELETE("/:id", handler.DeleteOrder)
}
