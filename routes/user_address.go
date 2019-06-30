package routes

import (
	"github.com/techforward/ECscript_server/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// UserAddressRouter User's Address Router
func UserAddressRouter(e *echo.Echo) {
	loginRequired := e.Group("/user_address")
	loginRequired.Use(middleware.JWT([]byte("secret")))

	loginRequired.GET("/", handler.GetAllUserAddresses)
	loginRequired.GET("/:id", handler.GetUserAddress)
	loginRequired.POST("/", handler.CreateUserAddress)
	loginRequired.PUT("/:id", handler.UpdateUserAddress)
	loginRequired.DELETE("/:id", handler.DeleteUserAddress)
}
