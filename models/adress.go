package models

import (
	"time"
)

type Address struct {
	Ulid     string `json:"ulid" example:"0000XSNJG0MQJHBF4QX1EFD6Y3" gorm:"primary_key"`
	Postcode string `json:"postcode" example:"163-0011"`
	Address1 string `json:"address1" example:"Kita-ku, Tokyo, Jp"`
	Address2 string `json:"address2" example:"23-3-12"`

	CreatedAt time.Time `json:"created_at" example:"2019-06-14 05:59:41"`
	UpdatedAt time.Time `json:"updated_at" example:"2019-06-15 15:01:52"`
}

// Validation Address
func (a *Address) Validation() error {
	switch {
	case len(a.Ulid) == 0:
		return ErrorAddressUlidInvalid
	case len(a.Postcode) == 0:
		return ErrorPostcodeInvalid
	case len(a.Address1) == 0:
		return ErrorAddressInvalid
	case len(a.Address2) == 0:
		return ErrorAddressInvalid
	default:
		return nil
	}
}
