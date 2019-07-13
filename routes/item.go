package routes

import (
	"github.com/techforward/ECscript_server/handler"

	"github.com/labstack/echo"
)

// ItemRouter ItemRouter
func ItemRouter(e *echo.Echo) {
	g := e.Group("/item")

	g.GET("", handler.GetAllItems)
	g.GET("/:id", handler.GetItem)

	g.POST("", handler.CreateItem)
	g.PUT("/:id", handler.UpdateItem)
	g.DELETE("/:id", handler.DeleteItem)

	// loginRequired := e.Group("/item")
	// loginRequired.Use(middleware.JWT([]byte("secret")))
	// loginRequired.POST("/", handler.CreateItem)
	// loginRequired.PUT("/:id", handler.UpdateItem)
	// loginRequired.DELETE("/:id", handler.DeleteItem)

	image := g.Group("/:id/image")
	image.GET("/", handler.GetAllItemImages)

	image.POST("/", handler.CreateItemImage)
	image.PUT("/:imageUlid", handler.UpdateItemImage)
	image.DELETE("/:imageUlid", handler.DeleteItemImage)

	// imageLoginRequired := g.Group("/image")
	// imageLoginRequired.Use(middleware.JWT([]byte("secret")))
	// imageLoginRequired.POST("/", handler.CreateItemImage)
	// imageLoginRequired.PUT("/:imageUlid", handler.UpdateItemImage)
	// imageLoginRequired.DELETE("/:imageUlid", handler.DeleteItemImage)
}
