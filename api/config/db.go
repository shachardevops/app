package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func init() {
	var err error
	pgUser := os.Getenv("POSTGRES_USER")
	pgPass := os.Getenv("POSTGRES_PASSWORD")
	pgDb := os.Getenv("POSTGRES_DB")
	data := fmt.Sprintf("postgres://%v:%v@postgres/%v?sslmode=disable", pgUser, pgPass, pgDb)
	DB, err = sql.Open("postgres", data)
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
