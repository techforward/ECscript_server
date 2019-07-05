package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/techforward/ECscript_server/config"
)

func ConnectGORM() *gorm.DB {
	config := config.SetEnvironment()

	DBMS := "mysql"
	USER := config.User
	PASS := config.Password
	IP := config.IP
	PORT := config.Port
	PROTOCOL := "tcp(" + IP + ":" + PORT + ")"
	DBNAME := config.Name
	PARSETIME := "parseTime=true"
	//TIMEZONE := "loc=Asia%2FTokyo"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?" + PARSETIME //+ "&" + TIMEZONE
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}
