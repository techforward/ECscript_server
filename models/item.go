package models

import (
	"time"
)

// Item ItemTable
type Item struct {
	Ulid     string `json:"ulid" example:"0000XSNJG0MQJHBF4QX1EFD6Y3" gorm:"primary_key"`
	Name     string `json:"name" example:"item name"`
	Context  string `json:"context" example:"item context"`
	Amount   int    `json:"amount" example:"4" format:"int64"`
	Priority int    `json:"priority" example:"1" format:"int64"`

	// CreatedAt, UpdatedAt, DeleteAt という time.Time 型のメンバを追加すると、
	// GORM側でここに追加・更新・削除日時を自動的に入るようになる。
	CreatedAt time.Time `json:"created_at" example:"2019-06-14 05:59:41"`
	UpdatedAt time.Time `json:"updated_at" example:"2019-06-15 15:01:52"`
	// DeleteAt  time.Time `json:"deleted_at" example:"2019-06-16 19:01:52"`
}

// Validation Item
func (a *Item) Validation() error {
	switch {
	case len(a.Ulid) == 0:
		return ErrorItemUlidInvalid
	case len(a.Name) == 0:
		return ErrorNameInvalid
	case len(a.Context) == 0:
		return ErrorContextInvalid
	case a.Amount < 0:
		return ErrorAmountInvalid
	case a.Priority < 0:
		return ErrorPriorityInvalid
	default:
		return nil
	}
}
