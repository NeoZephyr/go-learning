package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
)

var db *sql.DB

const (
	server = "cdh"
	port = 3306
	user = "root"
	password = "123456"
	database = "nba"
)

type hero struct {
	id int
	name string
	maxHp float32
}

func main() {
	var err error
	address := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, server, port, database)
	db, err = sql.Open("mysql", address)

	if err != nil {
		log.Fatalln(err.Error())
	}

	ctx := context.Background()
	err = db.PingContext(ctx)

	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Printf("connect to database %s success\n", database)

	hero, err := queryOne(10080)

	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println(hero)

	heroes, err := queryMany([]int{10000, 10001, 10002})

	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println(heroes)

	hero.name = "杜贝壳"
	update(hero)
	hero, _ = queryOne(10080)
	fmt.Println(hero)
}

func queryOne(id int) (hero, error) {
	hero := hero{}
	err := db.QueryRow("select id, name, hp_max from heros where id = ?", id).Scan(
		&hero.id, &hero.name, &hero.maxHp)
	return hero, err
}

func queryMany(ids []int) ([]hero, error) {
	args := make([]interface{}, len(ids))
	for i, id := range ids {
		args[i] = id
	}
	heroes := make([]hero, 0)
	sqlText := fmt.Sprintf("select id, name, hp_max from heros where id in (?%s)", strings.Repeat(", ?", len(ids) -1))
	println(sqlText)
	rows, err := db.Query(sqlText, args...)

	if err != nil {
		return heroes, err
	}

	for rows.Next() {
		hero := hero{}
		err := rows.Scan(&hero.id, &hero.name, &hero.maxHp)
		if err != nil {
			log.Fatalln(err.Error())
		}
		heroes = append(heroes, hero)
	}
	return heroes, err
}

func update(hero hero) {
	_, err := db.Exec("update heros set name = ?, hp_max = ? where id = ?", hero.name, hero.maxHp, hero.id)

	if err != nil {
		log.Fatalln(err.Error())
	}
	return
}

func delete(id int) {
	_, err := db.Exec("delete from heros where id = ?", id)

	if err != nil {
		log.Fatalln(err.Error())
	}

	return
}