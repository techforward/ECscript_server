package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Item struct {
	Ulid       string `json:"ulid" example:"0000XSNJG0MQJHBF4QX1EFD6Y3" `
	Name       string `json:"name" example:"item name"`
	Context    string `json:"context" example:"item context"`
	Amount     int    `json:"amount" example:"4" format:"int64"`
	Priority   int    `json:"priority" example:"1" format:"int64"`
	Created_at string `json:"created_at" example:"2019.."`
	Ureated_at string `json:"updated_at" example:"2019.."`
}

func ConnectGORM() *gorm.DB {
	DBMS := "mysql"
	USER := "api"
	PASS := "api"
	PROTOCOL := "tcp(35.221.187.41:3306)"
	DBNAME := "ecsite"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}

// func main() {
// 	var items []Item

// 	db := db.ConnectGORM()
// 	db.SingularTable(true)
// 	db.Find(&items)
// 	fmt.Println(items)
// }
