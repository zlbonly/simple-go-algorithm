package main

/*
linux 安装 配置kafka (安装kafka前确保安装了java)


1、安装
# 版本：kafka_2.11-2.3.0.tgz
wget http://mirrors.tuna.tsinghua.edu.cn/apache/kafka/2.3.0/kafka_2.11-2.3.0.tgz

tar -zxvf kafka_2.11-2.3.0.tgz

cp kafka_2.11-2.3.0 /usr/local/kafka
2、启动

cd /usr/local/kafka/bin

先启动zk  然后启动kafka
./zookeeper-server-start.sh -daemon ../config/zookeeper.properties

./kafka-server-start.sh -daemon ../config/server.properties


测试
1、创建一个topic
# 创建名为test的chart，只有一个副本，一个分区
kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic chart

# 查看kafka的topic
kafka-topics.sh -list -zookeeper localhost:2181

2、生产、消费测试
# 启动生产端
./kafka-console-producer.sh --broker-list localhost:9092 --topic chart

#启动消费端
./kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic chart --from-beginning
*/

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
