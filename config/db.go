package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("postgres", "postgres://bond:password@postgres/bookstore?sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}
	_, err = DB.Exec(`CREATE TABLE books (
		isbn    char(14)     PRIMARY KEY NOT NULL,
		title   varchar(255) NOT NULL,
		author  varchar(255) NOT NULL,
		price   decimal(5,2) NOT NULL
	  );`)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("You connected to your database.")
}
