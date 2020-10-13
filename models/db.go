package models

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func GoDotEnvVariable(key string) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}

	return os.Getenv(key), nil
}

// obtains db vars from .env file
func GetVarsDB() (string, string, string, string, error) {
	userDB, err := GoDotEnvVariable("USERDB")
	if err != nil {

		return "", "", "", "", err
	}

	passDB, err := GoDotEnvVariable("PASSDB")
	if err != nil {
		return "", "", "", "", err
	}

	ipDB, err := GoDotEnvVariable("IPDB")
	if err != nil {
		return "", "", "", "", err
	}

	nameDB, err := GoDotEnvVariable("NAMEDB")
	if err != nil {
		return "", "", "", "", err
	}

	return userDB, passDB, ipDB, nameDB, nil
}

// connect to database
func ConnectToBDD() {
	userDB, passDB, ipDB, nameDB, err := GetVarsDB()
	if err != nil {
		return
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", userDB, passDB, ipDB, nameDB)
	DB, _ = sql.Open("mysql", dsn)
}
