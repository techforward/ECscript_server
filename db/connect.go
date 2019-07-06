package db

import (
	"flag"
	"os"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/techforward/ECscript_server/config"
)

func ConnectGORM() *gorm.DB {
	var mode string

	if flag.Lookup("test.v") == nil {
		flag.StringVar(&mode, "mode", "test", "run mode")
		flag.VisitAll(func(f *flag.Flag) {
			if s := os.Getenv(strings.ToUpper(f.Name)); s != "" {
				f.Value.Set(s)
			}
		})
		flag.Parse()
	} else {
		mode = "test"
	}

	config := config.SetEnvironment(mode)

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
