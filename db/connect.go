package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/techforward/ECscript_server/config"
)

// ConnectGORM ConnectGORM
func ConnectGORM() *gorm.DB {
	var mode string

	// if os.Getenv("MODE") != "production" {

	// 	if flag.Lookup("test.v") == nil {
	// 		flag.StringVar(&mode, "mode", "test", "run mode")
	// 		flag.VisitAll(func(f *flag.Flag) {
	// 			if s := os.Getenv(strings.ToUpper(f.Name)); s != "" {
	// 				f.Value.Set(s)
	// 			}
	// 		})
	// 		flag.Parse()
	// 	} else {
	// 		mode = "test"
	// 	}

	// } else {
	mode = os.Getenv("MODE")
	// }
	config := config.SetEnvironment(mode)

	DBMS := "mysql"
	USER := config.User
	PASS := config.Password
	IP := config.IP
	PORT := config.Port
	DBNAME := config.Name
	//TIMEZONE := "loc=Asia%2FTokyo"

	var PARSETIME string
	var PROTOCOL string
	if os.Getenv("MODE") != "production" {
		PROTOCOL = "tcp(" + IP + ":" + PORT + ")"
		PARSETIME = "?parseTime=true"
	} else {
		// Cloud SQL
		PROTOCOL = "unix(/cloudsql/" + IP + ")"
		PARSETIME = "?parseTime=true"
	}

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + PARSETIME //+ "&" + TIMEZONE
	fmt.Println("CONNECT: " + CONNECT)

	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}
