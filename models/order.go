package models

import (
	"time"
)

type Order struct {
	Ulid        string `json:"ulid" example:"0000XSNJG0MQJHBF4QX1EFD6Y3" gorm:"primary_key"`
	UserUlid    string `json:"userUlid" example:"0000XSNJG0JFGRHF4QX1EFD6Y3"`
	AddressUlid string `json:"addressUlid" example:"0000XSNJG0JFGRHF4QX1EFDAH2"`
	IsCanceled  int    `json:"isCanceled" example:"0"`
	Status      int    `json:"status" example:"0"`

	CreatedAt time.Time `json:"created_at" example:"2019-06-14 05:59:41"`
	UpdatedAt time.Time `json:"updated_at" example:"2019-06-15 15:01:52"`
}

// Validation Order
func (a *Order) Validation() error {
	switch {
	case len(a.Ulid) == 0:
		return ErrorOrderUlidInvalid
	case len(a.UserUlid) == 0:
		return ErrorUserUlidInvalid
	case len(a.AddressUlid) == 0:
		return ErrorAddressUlidInvalid
	case a.IsCanceled < 0 || 1 < a.IsCanceled:
		return ErrorIsCanceledInvalid
	case a.Status < 0:
		return ErrorStatusInvalid
	default:
		return nil
	}
}
