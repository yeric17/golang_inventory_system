package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	User, Password, Host, Port, DBname, Driver, APIPort string
)

func init() {
	var err error
	err = godotenv.Load()
	if err != nil {
		panic(err)
	}

	User = os.Getenv("DB_USER")
	Password = os.Getenv("DB_PASS")
	Host = os.Getenv("DB_HOST")
	Port = os.Getenv("DB_PORT")
	DBname = os.Getenv("DB_NAME")
	Driver = os.Getenv("DB_DRIVER")
	APIPort = os.Getenv("API_PORT")
}
