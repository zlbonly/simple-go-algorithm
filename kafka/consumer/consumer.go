package main

import (
	"context"
	"flag"
	"github.com/segmentio/kafka-go"
	"log"
)

const (
	kafkaConn1 = "127.0.0.1:9092"
)

var (
	topic = flag.String("t", "test_kafka", "kafka_topic")
	group = flag.String("g", "test-group", "kafka consumer group")
)

func main() {
	flag.Parse()

	config := kafka.ReaderConfig{
		Brokers:  []string{kafkaConn1},
		GroupID:  *group,
		Topic:    *topic,
		MinBytes: 1e3,
		MaxBytes: 1e6,
	}

	reader := kafka.NewReader(config)

	ctx := context.Background()

	for {
		msg, err := reader.FetchMessage(ctx)

		if err != nil {
			log.Printf("fail to get msg:%v", err)
			continue
		}

		log.Printf("msg content:topic=%v,partition= %v,offset=%v,content=%v", msg.Topic, msg.Partition, msg.Offset, string(msg.Value))
		err = reader.CommitMessages(ctx, msg)

		if err != nil {
			log.Printf("fail to commit msg:%v", err)
		}
	}

}
