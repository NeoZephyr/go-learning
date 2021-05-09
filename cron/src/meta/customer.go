package meta

import (
	"cron/src/mysql"
	"encoding/json"
	"fmt"
)

type Customer struct {
	IsMember bool
	Age int
	Name string
}

func Get() {
	stmt, err := mysql.DBConnection().Prepare("select age, name from customer where id = ? limit 1")

	if err != nil {
		fmt.Printf("failed to prepare statement, error: %s\n", err.Error())
		return
	}

	defer stmt.Close()

	customer := Customer{}
	err = stmt.QueryRow(1).Scan( &customer.Age, &customer.Name)

	if err != nil {
		fmt.Printf("query failed, error: %s\n", err.Error())
		return
	}

	data, err := json.Marshal(customer)

	if err != nil {
		fmt.Printf("mashal customer failed, error: %s\n", err.Error())
		return
	}

	fmt.Printf("customer: %s\n", string(data))
}

func Store() {
	stmt, err := mysql.DBConnection().
		Prepare("insert ignore into customer(`is_member`, `age`, `name`) values (?, ?, ?)")

	if err != nil {
		fmt.Printf("failed to prepare statement, error: %s\n", err.Error())
		return
	}

	defer stmt.Close()

	ret, err := stmt.Exec(true, 10, "jack")

	if err != nil {
		fmt.Printf("failed to exec sql, error: %s\n", err.Error())
		return
	}

	if rows, err := ret.RowsAffected(); err == nil {
		if rows <= 0 {
			fmt.Printf("already insert")
		}
	}
}
