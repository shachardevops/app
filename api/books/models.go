package books

import (
	"app/api/config"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	Isbn   string  `json:"Isbn"`
	Title  string  `json:"Title"`
	Author string  `json:"Author"`
	Price  float64 `json:"Price"`
}

func AllBooks() ([]Book, error) {
	rows, err := config.DB.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bks := make([]Book, 0)
	for rows.Next() {
		bk := Book{}
		err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
		if err != nil {
			return nil, err
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return bks, nil
}

func OneBook(r *http.Request) (Book, error) {
	bk := Book{}
	isbn := mux.Vars(r)["id"]
	if isbn == "" {
		return bk, errors.New("400. Bad Request.")
	}

	row := config.DB.QueryRow("SELECT * FROM books WHERE isbn = $1", isbn)

	err := row.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
	if err != nil {
		return bk, err
	}

	return bk, nil
}

func PutBook(r *http.Request) (Book, error) {
	bk := Book{}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return bk, errors.New("Error." + err.Error())

	}
	json.Unmarshal(reqBody, &bk)
	_, err = config.DB.Exec("INSERT INTO books (isbn, title, author, price) VALUES ($1, $2, $3, $4)", bk.Isbn, bk.Title, bk.Author, bk.Price)
	if err != nil {
		return bk, errors.New("500. Internal Server Error." + err.Error())
	}
	return bk, nil
}

func UpdateBook(r *http.Request) (Book, error) {
	bk := Book{}
	isbn := mux.Vars(r)["id"]
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return bk, errors.New("400. Bad Request. Fields can't be empty.")
	}
	json.Unmarshal(reqBody, &bk)
	_, err = config.DB.Exec("UPDATE books SET isbn = $1, title=$2, author=$3, price=$4 WHERE isbn=$1;", isbn, bk.Title, bk.Author, bk.Price)
	if err != nil {
		return bk, err
	}

	return bk, nil
}

func DeleteBook(r *http.Request) error {
	isbn := mux.Vars(r)["id"]
	if isbn == "" {
		return errors.New("400. Bad Request.")
	}

	output, err := config.DB.Exec("DELETE FROM books WHERE isbn=$1;", isbn)
	fmt.Println(output)

	fmt.Println(err)
	if err != nil {
		return errors.New("500. Internal Server Error")
	}
	return nil
}
