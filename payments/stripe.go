package payments

import (
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"github.com/stripe/stripe-go/customer"
)

// CreateCustomer サーバーにはcustomerIDを保存 それを利用
func CreateCustomer(email string) (string, error) {
	// Set your secret key: remember to change this to your live secret key in production
	// See your keys here: https://dashboard.stripe.com/account/apikeys
	stripe.Key = "sk_test_4eC39HqLyjWDarjtT1zdp7dc"

	params := &stripe.CustomerParams{
		Email: stripe.String(email),
	}

	cus, err := customer.New(params)
	if err != nil {
		return "", err
	}

	cusID := cus.ID

	return cusID, nil
}

// CreateCharge 決済する
func CreateCharge(orderID string, customerID string, price int64, token string) (*stripe.Charge, error) {
	stripe.Key = "sk_test_4eC39HqLyjWDarjtT1zdp7dc"

	chargeParams := &stripe.ChargeParams{
		Customer:    stripe.String(customerID),
		Amount:      stripe.Int64(price),
		Currency:    stripe.String(string(stripe.CurrencyJPY)),
		Description: stripe.String("orderID is " + orderID),
	}
	chargeParams.SetSource(token) // obtained with Stripe.js
	ch, err := charge.New(chargeParams)
	if err != nil {
		return nil, err
	}
	return ch, nil
}
