package main

import (
	"exmaple/channelmq/pre_test/channel_mq"
	"fmt"
	"time"
)

var (
	topic = "Goland"
)

func main() {
	OnceTopic()
}

func OnceTopic() {

	m := channel_mq.NewClient()
	m.SetConditions(10)

	ch, err := m.Subscribe(topic)

	if err != nil {
		fmt.Print("subscribe failed")
	}

	go OncePub(m)

	OnceSub(ch, m)
	defer m.Close()

}

func OncePub(c *channel_mq.Client) {
	t := time.NewTicker(10 * time.Second)

	defer t.Stop()

	for {
		select {
		case <-t.C:
			err := c.Publish(topic, "真帅")
			if err != nil {
				fmt.Println("pub message failed")
			}
		default:

		}
	}
}

// 接受订阅消息
func OnceSub(m <-chan interface{}, c *channel_mq.Client) {
	for {
		val := c.GetPayLoad(m)
		fmt.Printf("get message is %s\n", val)
	}
}
