package models

import (
	"time"
)

// CartItem CartItemTable
type CartItem struct {
	Ulid      string `json:"ulid" example:"0000XSNJG0MQJHBF4QX1EFD6Y3" gorm:"primary_key"`
	CartUlid  string `json:"order_ulid" example:"0000XSNJG0JFGRHF4QX1EFD6Y3"`
	ItemUlid  string `json:"item_ulid" example:"0000XSNJG0JFGRHF4QX1EFDAH2"`
	Amount    int    `json:"amount" example:"10"`
	IsDeleted int    `json:"is_deleted" example:"0"`
	IsKeep    int    `json:"is_keep" example:"0"`

	CreatedAt time.Time `json:"created_at" example:"2019-06-14 05:59:41"`
	UpdatedAt time.Time `json:"updated_at" example:"2019-06-15 15:01:52"`
}

// Validation CartItem
func (a *CartItem) Validation() error {
	switch {
	case len(a.Ulid) == 0:
		return ErrorCartItemUlidInvalid
	case len(a.CartUlid) == 0:
		return ErrorCartUlidInvalid
	case len(a.ItemUlid) == 0:
		return ErrorItemUlidInvalid
	case a.Amount < 0:
		return ErrorAmountInvalid
	case a.IsDeleted < 0:
		return ErrorIsDeletedInvalid
	case a.IsKeep < 0:
		return ErrorIsKeepInvalid
	default:
		return nil
	}
}
