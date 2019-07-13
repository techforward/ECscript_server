package models

import (
	"time"
)

// Other OtherTable
type Other struct {
	Ulid string `json:"ulid" example:"0000XSNJG0MQJHBF4QX1EFD6Y3" gorm:"primary_key"`
	Name string `json:"name" example:"item name"`
	Path string `json:"path" example:"item path"`

	CreatedAt time.Time `json:"created_at" example:"2019-06-14 05:59:41"`
	UpdatedAt time.Time `json:"updated_at" example:"2019-06-15 15:01:52"`
}
