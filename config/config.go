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

func setDatabase(mode string) Database {
	return Database{
		User:     os.Getenv("DB_USER" + mode),
		Password: os.Getenv("DB_PASS" + mode),
		Name:     os.Getenv("DB_PASS" + mode),
		IP:       os.Getenv("DB_PASS" + mode),
		Port:     os.Getenv("DB_PASS" + mode),
	}

}

// SetEnvironment Set nvironment
func SetEnvironment() Database {
	var database Database
	var mode string

	flag.StringVar(&mode, "mode", "", "run mode")
	flag.VisitAll(func(f *flag.Flag) {
		if s := os.Getenv(strings.ToUpper(f.Name)); s != "" {
			f.Value.Set(s)
		}
	})
	flag.Parse()

	switch mode {
	case "production":
		database = setDatabase("")
	case "local":
		database = setDatabase("_LOCAL")
	default:
		database = setDatabase("_TEST")
	}

	return database
}
