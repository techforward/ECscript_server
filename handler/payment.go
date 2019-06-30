package handler

import (
	"github.com/techforward/ECscript_server/payments"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// StripeCharge godoc
// @Summary 決済
// @Description 決済を行う
// @Tags Stripe
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 404
// @Router /payment/stripe [post]
func StripeCharge(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	customerID := claims["stripeId"].(string)

	orderID := c.FormValue("orderId")
	price, _ := strconv.Atoi(c.FormValue("price"))
	token := c.FormValue("token")

	charge, err := payments.CreateCharge(orderID, customerID, int64(price), token)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, charge)
}
