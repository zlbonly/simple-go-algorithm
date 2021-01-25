package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
)

var (
	tcpNsqAddr = "127.0.0.1:4150"
)

func main() {
	config := nsq.NewConfig()
	for i := 0; i < 100; i++ {
		tPro, err := nsq.NewProducer(tcpNsqAddr, config)
		if err != nil {
			fmt.Println(err)
		}

		topic := "Insert"

		tCommand := "new data!"

		err = tPro.Publish(topic, []byte(tCommand))
		if err != nil {
			fmt.Println(err)
		}
	}
}
