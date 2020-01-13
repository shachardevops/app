package config

import (
	"log"

	"github.com/nsqio/go-nsq"
)

var Producer *nsq.Producer

func init() {
	var err error
	config := nsq.NewConfig()
	Producer, err = nsq.NewProducer("nsqd:4150", config)
	if err != nil {
		log.Panic(err)
	}
}
