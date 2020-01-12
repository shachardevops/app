package main

import (
	"app/books"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/nsqio/go-nsq"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	/*
		StrictSlash defines the slash behavior for new routes.
		When true, if the route path is "/path/", accessing "/path" will redirect to the former and vice versa.
	*/
	go consumer()
	router.HandleFunc("/books", books.Index)
	router.HandleFunc("/books/create", books.Create).Methods("POST")
	router.HandleFunc("/books/{id}", books.ShowOne).Methods("GET")
	router.HandleFunc("/books/{id}", books.Update).Methods("PATCH")
	router.HandleFunc("/books/{id}", books.Delete).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))

}

func consumer() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	decodeConfig := nsq.NewConfig()
	c, err := nsq.NewConsumer("api", "My_NSQ_Channel", decodeConfig)
	if err != nil {
		log.Panic("Could not create consumer")
	}
	//c.MaxInFlight defaults to 1

	c.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Println("NSQ message received:")
		log.Println(string(message.Body))
		return nil
	}))

	err = c.ConnectToNSQD("nsqd:4150")
	if err != nil {
		log.Panic("Could not connect")
	}
	log.Println("Awaiting messages from NSQ topic \"My NSQ Topic\"...")
	wg.Wait()
}
