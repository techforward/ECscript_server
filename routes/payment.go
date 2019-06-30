package routes

import (
	"github.com/techforward/ECscript_server/handler"

	"github.com/labstack/echo"
)

func PaymentRouter(e *echo.Echo) {
	g := e.Group("/payment")

	stripe := g.Group("/stripe")

	stripe.POST("/", handler.StripeCharge)
}
