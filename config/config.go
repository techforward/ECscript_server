package config

import (
	"flag"
	"os"
	"strings"
)

type Database struct {
	User     string
	Password string
	Name     string
	IP       string
	Port     string
}

func (db *Database) setDatabase(mode string) {
	db.User = os.Getenv("DB_USER" + mode)
	db.Password = os.Getenv("DB_PASS" + mode)
	db.Name = os.Getenv("DB_NAME" + mode)
	db.IP = os.Getenv("DB_ADDR" + mode)
	db.Port = os.Getenv("DB_PORT" + mode)
}

// SetEnvironment Set nvironment
func SetEnvironment() Database {
	var database Database
	var mode string

	flag.StringVar(&mode, "mode", "test", "run mode")
	flag.VisitAll(func(f *flag.Flag) {
		if s := os.Getenv(strings.ToUpper(f.Name)); s != "" {
			f.Value.Set(s)
		}
	})
	flag.Parse()

	switch mode {
	case "production":
		database.setDatabase("")
	case "local":
		database.setDatabase("_LOCAL")
	default:
		database.setDatabase("_TEST")
	}

	return database
}
