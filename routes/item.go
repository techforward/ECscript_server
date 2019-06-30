package routes

import (
	"github.com/techforward/ECscript_server/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func ItemRouter(e *echo.Echo) {
	g := e.Group("/item")

	g.GET("/", handler.GetAllItems)
	g.GET("/:id", handler.GetItem)

	loginRequired := e.Group("/item")
	loginRequired.Use(middleware.JWT([]byte("secret")))
	loginRequired.POST("/", handler.CreateItem)
	loginRequired.PUT("/:id", handler.UpdateItem)
	loginRequired.DELETE("/:id", handler.DeleteItem)

	image := g.Group("/image")
	image.GET("/", handler.GetAllItemImages)

	imageLoginRequired := g.Group("/image")
	imageLoginRequired.Use(middleware.JWT([]byte("secret")))
	imageLoginRequired.POST("/", handler.CreateItemImage)
	imageLoginRequired.PUT("/:imageUlid", handler.UpdateItemImage)
	imageLoginRequired.DELETE("/:imageUlid", handler.DeleteItemImage)
}
