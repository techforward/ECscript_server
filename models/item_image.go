package models

import (
	"time"
)

// ItemImage ItemImageTable
type ItemImage struct {
	Ulid     string `json:"ulid" example:"0000XSNJG0MQJHBF4QX1EFD6Y3" gorm:"primary_key"`
	ItemUlid string `json:"itemUlid" example:"0000XSNJG0JFGRHF4QX1EFD6Y3"`
	Path     string `json:"path" example:"asdada.png"`

	CreatedAt time.Time `json:"created_at" example:"2019-06-14 05:59:41"`
	UpdatedAt time.Time `json:"updated_at" example:"2019-06-15 15:01:52"`
}

// Validation ItemImage
func (a *ItemImage) Validation() error {
	switch {
	case len(a.Ulid) == 0:
		return ErrorItemImageUlidInvalid
	case len(a.ItemUlid) == 0:
		return ErrorItemUlidInvalid
	case len(a.Path) == 0:
		return ErrorPathInvalid
	default:
		return nil
	}
}
