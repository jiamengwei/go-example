package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@/go_blog")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	insertUser, err := db.Prepare("insert into user (username, password) values (?,?)")
	checkErr(err)
	res, err := insertUser.Exec("paul", "123456")
	checkErr(err)
	fmt.Println(res.LastInsertId())

	rows, err := db.Query("select * from user")
	for rows.Next() {
		var id int
		var username, password string
		err := rows.Scan(&id, &username, &password)
		checkErr(err)
		fmt.Println(id, username, password)
	}

}

type user struct {
	id       int
	username string
	password string
}
