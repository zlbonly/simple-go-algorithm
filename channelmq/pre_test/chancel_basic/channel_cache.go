package main

import (
	"sync"
	"time"
)

func main() {
	c := make(chan string, 2)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		c <- "Golang 梦工厂"
		c <- "asong"
	}()

	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 1)
		println("g公众号" + <-c)
		println("作者" + <-c)
	}()
}
