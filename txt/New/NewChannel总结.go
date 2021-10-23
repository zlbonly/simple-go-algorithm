package main

import (
	"fmt"
	"runtime"
	"sync"
)

/**
	Golang channel 专题。
	1、channel 原理
			1、我们使用make语句初始化一个channel时,将在当前程序堆上为一个hchan申请空间，创建一个hchan的结构体,并返回指向这个空间的指针。
	2、hchan结构体中的元素
		type hchan struct {
			qcount   uint   // 当前队列剩余的元素个数
			dataqsiz uint   // 环形队列的长度，即缓冲区的大小。
			buf      unsafe.Pointer
			elemsize uint16
			closed   uint32
			elemtype *_type // element type
			sendx    uint   // send index
			recvx    uint   // receive index
			recvq    waitq  // list of recv waiters
			sendq    waitq  // list of send waiters
			lock mutex
		}

		1、重点介绍的结构：
			qcount   uint   // 当前队列剩余的元素个数
			dataqsiz uint   // 环形队列的长度，即缓冲区的大小。
			buf      unsafe.Pointer // 环形队列的指针，用来存放缓冲区中的元素
			closed   uint32			// 当前通道的关闭状态
			recvq    waitq  // 等待读消息的goroutine队列
			sendq    waitq  // 等待写消息的goroutine队列
			lock mutex		// 互斥锁，为每个读写操作锁定通道，保护hchan并发安全性

		2、向channel写入数据流程
			1、锁定整个通道结构
			2、判断当前等待读取消息的recvq队列是否为空，如果不为空，说明缓冲区中没有数据，或者没有缓冲区，此时直接从
				recvq队列中取出G，并把数据写入，最后把该G唤醒，结束发送过程
			3、如果等待读取消息的recvq队列为空，判断缓冲区是否为空，如果缓冲区不为空，则向缓冲区写入数据，结束发送流程
				如果缓冲区为空，将数据写入当前的G，并把当前的G加入到sendq，然后进入睡眠，等待被读取数据的goroutine唤醒
			4、解除mutex锁

			流程图参考：
			https://imgconvert.csdnimg.cn/aHR0cHM6Ly9vc2NpbWcub3NjaGluYS5uZXQvb3NjbmV0L2M0YmE0MDEzMDE4MmJmNDI2NGFkNDU4YTJmMDU4NjNiZWYxLmpwZw

		3、从channel读取数据
			1、锁定整个通道结构
			2、判断当前等待发送数据的sendq队列是否为空，如果sendq不为空
				1、如果sendq不为空，且没有缓冲区，直接从sendq中取出G，把G中数据读出 ，最后把G唤醒，结束读取过程。
				2、如果sendq不为空，有缓冲区，则从buf环形队列首部取出一个元素，然后从sendq队列中取出一个G，将G中的数据写入buf环形队列尾部，结束流程。
			3、如果sendq队列为空，判断buf缓冲区是否有数据
				1 如果buf缓冲区有数据，直接读取缓冲区数据，结束流程。
				2、如果buf缓冲区没有数据或者缓冲区为空，将当前的Goroutine加入到recvq队列中，进入睡眠，等待被写gotourine唤醒,结束流程
			4、解除mutex锁

		4、关闭channel的流程
			1、释放recvq队列中所有的读取消息协程。接收操作可以正常运行，在缓存区中剩余数据读取完毕后，将读取到通道对应类型的零值。
			2、释放所有的sendq队列中的协程。但是要注意如果关闭后，继续写入会导致panic。

	2、channel 阻塞和panic 情况
			2.1	channel产生的panic
				1、关闭1个nil值的channel会引发panic  （var c chan int）
				2、关闭一个已关闭的channel会引发panic
				3、向一个已关闭的channel发送数据

			2.2 读写阻塞情况
				1、读写nil channel 会永久阻塞
				2、无缓冲通道阻塞
					1、通道中无数据，但执行读通道。
					2、通道中无数据，向通道写数据，但无协程读取。
				3、有缓存通道阻塞
					1、通道的缓存无数据，但执行读通道。
					2、通道的缓存已经占满，向通道写数据，但无协程读。
			2.3 无缓存通道 和缓存为1的通道的区别
				c := make(chan int) 与 c := make(chan int, 1) 的区别。

				1、无缓冲的channel的读写者必须同时完成发送和接收，而不能串行，
					显然单协程无法满足。所以这里造成了循环等待，会死锁。
					func main() {
						ch := make(chan int)
						ch <- 1
       					 <- ch
					}
				2、有缓冲的通道并不强制channel的读写者必须同时完成发送和接收，
 					读者只会在没有数据时阻塞，写者只会在没有可用容量时阻塞

				ch := make(chan int, 1)
				ch <- 1
				fmt.Println(<-ch)

			2.4 channel导致死锁
				1、同一个goroutine中，使用同一个 channel 读写。
						func main(){
							ch:=make(chan int)  //这就是在main程里面发生的死锁情况 (有缓存通道和无缓存通道都会导致)
							ch<-6   //  这里会发生一直阻塞的情况，执行不到下面一句
							<-ch
						}
				 2、第二种：2个 以上的go程中， 使用同一个 channel 通信。 读写channel 先于 go程创建。
					func main(){
						ch:=make(chan int)
						ch<-666    //这里一直阻塞，运行不到下面
						go func (){
							<-ch  //这里虽然创建了子go程用来读出数据，但是上面会一直阻塞运行不到下面
						}()
					}

	3、channel 场景编程题
		1、使用两个channel按顺序打印两个数组
		例如  str := []string{1,2,3,4},str2 := []string{A,B,C,D}按顺序打印 1A2B3C4D
			printOrderChannel()

		2、按顺序交替打印1-100的奇数和偶数
		printOddAndEven()

		3、生产消费模型。生产者 随机生成 1 ～ 30 数字。消费者开启两个协程消费生产者生产的数据。
		printProductAndConsumer()

		4、golang中使用channel实现互斥锁
		channelMutex()
*/

/**
golang中使用channel实现互斥锁
1、使用channel实现mutex锁
*/
type MyLock struct {
	ch chan struct{}
}

func (t *MyLock) Lock() {
	<-t.ch
}
func (t *MyLock) UnLock() {
	t.ch <- struct{}{}
}

func NewMyLock() *MyLock {
	m := &MyLock{
		ch: make(chan struct{}, 1),
	}
	m.ch <- struct{}{}
	return m
}

var counter int

func add(n int) {
	tmp := counter
	tmp += n
	counter = tmp
}
func channelMutex() {
	m := NewMyLock()
	wg := sync.WaitGroup{}
	wg.Add(200)

	for i := 0; i < 200; i++ {
		go func(num int) {
			m.Lock()
			defer m.UnLock()
			add(num)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Printf("numGoroutine: %v \n", runtime.NumGoroutine())
	fmt.Printf("counter: %v \n", counter)
}

/**
1、生产消费模型。生产者 随机生成 1 ～ 30 数字。消费者开启两个协程消费生产者生产的数据。
*/

func printProductAndConsumer() {
	var wg sync.WaitGroup
	ch1 := make(chan int, 2)

	// 生产者生成数据
	go func() {
		defer func() {
			close(ch1)
			wg.Done()
		}()
		for i := 0; i < 30; i++ {
			fmt.Printf("producer:data:%v \n", i)
			ch1 <- i
		}
	}()

	for i := 0; i < 2; i++ {
		go func(consumer int) {
			defer wg.Done()
			for {
				select {
				case value, ok := <-ch1:
					{
						if !ok {
							return
						}
						fmt.Printf("comsumer:%v, data:%v \n", consumer, value)
					}
				}
			}
		}(i)
	}
	wg.Add(3)
	wg.Wait()
}

// channel场景题，
/**
1、使用两个channel按顺序打印两个数组
例如  str := []string{1,2,3,4},str2 := []string{A,B,C,D}按顺序打印 1A2B3C4D
*/

func printOrderChannel() {
	str1 := []string{"1", "2", "3", "4"}
	str2 := []string{"A", "B", "C", "D"}
	ch1 := make(chan bool, 1)
	ch2 := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for _, value1 := range str1 {
			if ok := <-ch1; ok {
				fmt.Printf("%v ", value1)
				ch2 <- true
			}
		}
	}()

	go func() {
		defer func() {
			close(ch1)
			close(ch2)
			wg.Done()
		}()
		for _, value2 := range str2 {
			if ok := <-ch2; ok {
				fmt.Printf("%v ", value2)
				ch1 <- true
			}
		}
	}()

	ch1 <- true
	wg.Wait()
}

/**
1、按顺序交替打印1-100的奇数和偶数
*/
func printOddAndEven() {
	var wg sync.WaitGroup
	ch1 := make(chan bool, 1)
	ch2 := make(chan bool)
	wg.Add(2)
	// 打印奇数
	go func() {
		defer wg.Done()
		for i := 1; i < 50; i++ {
			if ok := <-ch1; ok {
				fmt.Printf("%v ", i*2-1)
				ch2 <- true
			}
		}
	}()

	// 打印偶数
	go func() {
		defer func() {
			close(ch1)
			close(ch2)
			wg.Done()
		}()
		for i := 1; i < 50; i++ {
			if ok := <-ch2; ok {
				fmt.Printf("%v ", i*2)
				ch1 <- true
			}
		}
	}()
	ch1 <- true
	wg.Wait()
}
