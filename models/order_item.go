package models

import (
	"time"
)

type OrderItem struct {
	Ulid      string `json:"ulid" example:"0000XSNJG0MQJHBF4QX1EFD6Y3" gorm:"primary_key"`
	OrderUlid string `json:"order_ulid" example:"0000XSNJG0JFGRHF4QX1EFD6Y3"`
	ItemUlid  string `json:"item_ulid" example:"0000XSNJG0JFGRHF4QX1EFDAH2"`
	Amount    int    `json:"amount" example:"10"`
	Price     int    `json:"price" example:"224"`

	CreatedAt time.Time `json:"created_at" example:"2019-06-14 05:59:41"`
	UpdatedAt time.Time `json:"updated_at" example:"2019-06-15 15:01:52"`
}

// Validation OrderItem
func (a *OrderItem) Validation() error {
	switch {
	case len(a.Ulid) == 0:
		return ErrorOrderItemUlidInvalid
	case len(a.OrderUlid) == 0:
		return ErrorOrderUlidInvalid
	case len(a.ItemUlid) == 0:
		return ErrorItemUlidInvalid
	case a.Amount < 0:
		return ErrorAmountInvalid
	case a.Price < 0:
		return ErrorPriceInvalid
	default:
		return nil
	}
}
