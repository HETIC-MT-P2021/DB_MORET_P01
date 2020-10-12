package models

import (
	"database/sql"
	"fmt"
	_ "fmt"
	_ "github.com/go-sql-driver/mysql"
)

var UserDb = "xxx"
var PassDb = "xxx"
var IpDb = "db:3306"
var NameDb = "classicmodels"

var DB *sql.DB

func ConnectDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", UserDb, PassDb, IpDb, NameDb)
	DB, _ = sql.Open("mysql", dsn)
}