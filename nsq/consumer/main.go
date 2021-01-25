package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"sync"
	"time"
)

var (
	tcpNsqdAddr = "127.0.0.1:4150"
)

type NsqHandler struct {
	msqCount     int64
	nsqHandlerID string
}

func (s *NsqHandler) HandleMessage(message *nsq.Message) error {

	s.msqCount++

	fmt.Println(s.msqCount, s.nsqHandlerID)
	fmt.Printf("msg.Timestamp=%v,msg.nsqaddress=%s,msg.body=%s\n", time.Unix(0, message.Timestamp).Format("2021-01-23 10:00:00"), message.NSQDAddress, string(message.Body))
	return nil
}

func main() {
	config := nsq.NewConfig()
	com, err := nsq.NewConsumer("Insert", "channel1", config)
	if err != nil {
		fmt.Println(err)
	}

	com.AddHandler(&NsqHandler{nsqHandlerID: "one"})

	err = com.ConnectToNSQD(tcpNsqdAddr)
	if err != nil {
		fmt.Println(err)
	}

	var wg = &sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
