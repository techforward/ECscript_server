package models

import (
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	// gorm:"primary_key"を消すとテーブルが空っぽになる
	Ulid        string  `json:"ulid" example:"0000XSNJG0MQJHBF4QX1EFD6Y3" gorm:"primary_key"`
	FirebaseUID *string `json:"firebaseUid" example:"0gwEMgeG0ePdEZOZV6ZchBgWyE42"`
	StripeID    *string `json:"stripeId" example:"cus_fherfjwh234rhej"`
	CartUlid    *string `json:"cartUlid" example:"0000XSHJW0MQJHBF4QX1EFD6Y3"`
	AddressUlid *string `json:"addressUlid" example:"0000XSHJW0MQJ3JK3QX1EFD6Y3"`
	Name        string  `json:"name" example:"user124"`
	Email       string  `json:"email" example:"user@gmail.com"`

	CreatedAt time.Time `json:"created_at" example:"2019-06-14 05:59:41"`
	UpdatedAt time.Time `json:"updated_at" example:"2019-06-15 15:01:52"`
}

// Validation User
func (a *User) Validation() error {
	switch {
	case len(a.Ulid) == 0:
		return ErrorUserUlidInvalid
	case len(a.Name) == 0:
		return ErrorNameInvalid
	case checkmail.ValidateFormat(a.Email) != nil:
		return ErrorEmailInvalid
	// case len(a.CartUlid) == 0:
	// 	return ErrorCartUlidInvalid
	default:
		return nil
	}
}
