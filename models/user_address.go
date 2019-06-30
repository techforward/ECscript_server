package models

import (
	"time"
)

type UserAddress struct {
	Ulid        string `json:"ulid" example:"0000XSNJG0MQJHBF4QX1EFD6Y3" gorm:"primary_key"`
	UserUlid    string `json:"user_ulid" example:"0000XSNJG0JFGRHF4QX1EFD6Y3"`
	AddressUlid string `json:"address_ulid" example:"0000XSNJG0JFGRHF4QX1EFDAH2"`

	CreatedAt time.Time `json:"created_at" example:"2019-06-14 05:59:41"`
	UpdatedAt time.Time `json:"updated_at" example:"2019-06-15 15:01:52"`
}

// Validation UserAddress
func (a *UserAddress) Validation() error {
	switch {
	case len(a.Ulid) == 0:
		return ErrorUserAddressUlidInvalid
	case len(a.UserUlid) == 0:
		return ErrorUserUlidInvalid
	case len(a.AddressUlid) == 0:
		return ErrorAddressUlidInvalid
	default:
		return nil
	}
}
