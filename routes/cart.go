package routes

import (
	"github.com/techforward/ECscript_server/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func CartRouter(e *echo.Echo) {
	g := e.Group("/cart")

	g.GET("/", handler.GetAllCarts)
	g.GET("/:id", handler.GetCart)

	loginRequired := e.Group("/cart")
	loginRequired.Use(middleware.JWT([]byte("secret")))
	loginRequired.POST("/", handler.CreateCart)
	loginRequired.PUT("/:id", handler.UpdateCart)
	loginRequired.DELETE("/:id", handler.DeleteCart)
}
