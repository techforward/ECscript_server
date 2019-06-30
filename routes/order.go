package routes

import (
	"github.com/techforward/ECscript_server/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func OrderRouter(e *echo.Echo) {
	g := e.Group("/order")

	g.GET("/", handler.GetAllOrders)
	g.GET("/:id", handler.GetOrder)

	loginRequired := e.Group("/order")
	loginRequired.Use(middleware.JWT([]byte("secret")))
	loginRequired.POST("/", handler.CreateOrder)
	loginRequired.PUT("/:id", handler.UpdateOrder)
	loginRequired.DELETE("/:id", handler.DeleteOrder)
}
