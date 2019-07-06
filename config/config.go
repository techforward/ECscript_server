package config

import (
	"os"
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
func SetEnvironment(mode string) Database {
	var database Database

	switch mode {
	case "production":
		database.setDatabase("")
	case "local":
		database.setDatabase("_LOCAL")
	case "test":
		database.setDatabase("_TEST")
	default:
		database.setDatabase("_TEST")
	}

	// fmt.Printf("%v\n", database)

	return database
}
