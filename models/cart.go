package models

import (
	"time"
)

// Cart CartTable
type Cart struct {
	Ulid     string     `json:"ulid" example:"0000XSNJG0MQJHBF4QX1EFD6Y3" gorm:"primary_key"`
	UserUlid string     `json:"userUlid" example:"0000XSNJG0JFGRHF4QX1EFD6Y3"`
	BoughtAt *time.Time `json:"boughtAt" example:"2019-06-15 15:59:41"`

	CreatedAt time.Time `json:"created_at" example:"2019-06-14 05:59:41"`
	UpdatedAt time.Time `json:"updated_at" example:"2019-06-15 15:01:52"`
}

// Validation Cart
func (a *Cart) Validation() error {
	switch {
	case len(a.Ulid) == 0:
		return ErrorCartUlidInvalid
	case len(a.UserUlid) == 0:
		return ErrorUserUlidInvalid
	// case a.BoughtAt.Unix() > time.Now().Unix():
	// 	return ErrorBoughtAtInvalid
	default:
		return nil
	}
}
