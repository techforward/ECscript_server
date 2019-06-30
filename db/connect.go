package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/techforward/ECscript_server/config"
)

func ConnectGORM() *gorm.DB {
	env := "local_development"
	if env == "local_test" {
		// mysqld, err := mysqltest.NewMysqld(nil)
		// if err != nil {
		// 	log.Fatalf("Failed to start mysqld: %s", err)
		// }
		// defer mysqld.Stop()

		// db, err := sql.Open("mysql", mysqld.Datasource("test", "", "", 0))

		// if err != nil {
		// 	panic(err.Error())
		// }
		// return db
	}
	config := config.SetEnvironment("local_development")

	DBMS := "mysql"
	USER := config.Database.User
	PASS := config.Database.Password
	IP := config.Database.IP
	PORT := config.Database.Port
	PROTOCOL := "tcp(" + IP + ":" + PORT + ")"
	DBNAME := config.Database.Name
	PARSETIME := "parseTime=true"
	//TIMEZONE := "loc=Asia%2FTokyo"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?" + PARSETIME //+ "&" + TIMEZONE
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}
