package books

import (
	"app/api/config"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func producer(t string, m string) {

	err := config.Producer.Publish(t, []byte(m))
	if err != nil {
		log.Panic(err)
	}
}
func Index(w http.ResponseWriter, r *http.Request) {
	bks, err := AllBooks()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	js, err := json.Marshal(bks)
	if err != nil {
		log.Println(err)
	}
	producer("api", "GET ALL")
	w.Write(js)
}

func ShowOne(w http.ResponseWriter, r *http.Request) {
	bk, err := OneBook(r)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	producer("api", "GET ONE")
	err = json.NewEncoder(w).Encode(bk)
	if err != nil {
		log.Println(err)
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	bk, err := PutBook(r)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusNotAcceptable)
		return
	}
	fmt.Println(bk)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	producer("api", "POST")
	json.NewEncoder(w).Encode(bk)
}

func Update(w http.ResponseWriter, r *http.Request) {
	bk, err := UpdateBook(r)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusBadRequest)
		return
	}
	fmt.Println(bk)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	producer("api", "UPDATE")
	json.NewEncoder(w).Encode(bk)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	err := DeleteBook(r)
	isbn := mux.Vars(r)["id"]
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	producer("api", "DELETE")
	fmt.Fprintf(w, "Mission accomplished: The book with the id %v has been deleted ", isbn)
}
