package models

import (
	"time"
)

type Price struct {
	Ulid     string `json:"ulid" example:"0000XSNJG0MQJHBF4QX1EFD6Y3" gorm:"primary_key"`
	ItemUlid string `json:"item_ulid" example:"0000XSNJG0JFGRHF4QX1EFD6Y3"`
	Price    int    `json:"price" example:"23"`

	CreatedAt time.Time `json:"created_at" example:"2019-06-14 05:59:41"`
	UpdatedAt time.Time `json:"updated_at" example:"2019-06-15 15:01:52"`
}

// Validation Price
func (a *Price) Validation() error {
	switch {
	case len(a.Ulid) == 0:
		return ErrorPriceUlidInvalid
	case len(a.ItemUlid) == 0:
		return ErrorNameInvalid
	case a.Price < 0:
		return ErrorPriceInvalid
	default:
		return nil
	}
}
