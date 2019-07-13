package routes

import (
	"github.com/labstack/echo"
)

// Router echo routers
func Router(e *echo.Echo) {
	AuthRouter(e)
	AddressRouter(e)
	CartRouter(e)
	ItemRouter(e)
	OrderRouter(e)
	UserRouter(e)
	UserAddressRouter(e)
	OtherRouter(e)
}
