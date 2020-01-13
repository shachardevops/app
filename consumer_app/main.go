package main

import (
	"fmt"
	"log"

	"github.com/nsqio/go-nsq"
)

func main() {
	messages := make(chan string)
	quit := make(chan string)
	defer close(quit)
	defer close(messages)
	go consumer(messages, quit)
	receive(messages, quit)

}

func receive(m, q <-chan string) {
	for {
		select {
		case v := <-m:
			fmt.Println(`Message from "message" channel:`, v)
		case i, ok := <-q:
			if !ok {
				fmt.Println("Quit from the channel:", i, ok)
				return
			}
			fmt.Println("From comma ok", i)

		}
	}
}

func consumer(m, q chan<- string) {
	decodeConfig := nsq.NewConfig()
	c, err := nsq.NewConsumer("api", "My_NSQ_Channel", decodeConfig)
	if err != nil {
		log.Panic("Could not create consumer")
	}

	c.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Println("NSQ message received:")
		log.Println(string(message.Body))
		m <- string(message.Body)
		return nil
	}))

	err = c.ConnectToNSQD("nsqd:4150")
	if err != nil {
		log.Panic("Could not connect")
	}
	log.Println("Awaiting messages from NSQ topic \"My NSQ Topic\"...")
}
