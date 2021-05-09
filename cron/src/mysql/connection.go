package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var Db *sql.DB

func init() {
	Db, _ = sql.Open("mysql", "root:123456@tcp(cdh:3306)/app?charset=UTF8")
	Db.SetMaxOpenConns(20)
	err := Db.Ping()

	if err != nil {
		fmt.Printf("failed to connect mysql, error: %s\n", err.Error())
		os.Exit(1)
	}
}

func DBConnection() *sql.DB {
	return Db
}

func ParseRows(rows *sql.Rows) []map[string]interface{} {
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for j := range values {
		scanArgs[j] = &values[j]
	}

	record := make(map[string]interface{})
	records := make([]map[string]interface{}, 0)
	for rows.Next() {
		err := rows.Scan(scanArgs...)
		checkErr(err)

		for i, col := range values {
			if col != nil {
				record[columns[i]] = col
			}
		}
		records = append(records, record)
	}
	return records
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}