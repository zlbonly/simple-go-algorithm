package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"github.com/segmentio/kafka-go"
	"os"
)

const (
	kafkaConn1 = "127.0.0.1:9092"
	topic      = "test_kafka"
)

var brokerAddrs = []string{kafkaConn1}

func newKafkaWriter() *kafka.Writer {
	return &kafka.Writer{
		Addr:  kafka.TCP(kafkaConn1),
		Topic: topic,
	}
}
func main() {
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	writer := newKafkaWriter()
	defer writer.Close()

	for {
		fmt.Print("Enter msg...")
		msgStr, _ := reader.ReadString('\n')

		msg := kafka.Message{
			Value: []byte(msgStr),
		}

		err := writer.WriteMessages(context.Background(), msg)
		if err != nil {
			fmt.Println(err)
		}

	}
}
