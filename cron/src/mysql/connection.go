package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "root:123456@tcp(cdh:3306)/app?charset=UTF8")
	db.SetMaxOpenConns(20)
	err := db.Ping()

	if err != nil {
		fmt.Printf("failed to connect mysql, error: %s\n", err.Error())
		os.Exit(1)
	}
}

func DBConnection() *sql.DB {
	return db
}