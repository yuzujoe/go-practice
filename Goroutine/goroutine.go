package main

import (
	"fmt"
	"sync"
	"time"
)

// sync.WaitGroup
func goroutine(s string, wg *sync.WaitGroup) {
	// Doneを呼ばないとエラーになる
	defer wg.Done()
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func normal(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func helloworld() {
	var wg sync.WaitGroup
	wg.Add(1)
	// これで並列処理ができる
	go goroutine("hello", &wg)
	normal("World")
	wg.Wait()
}

// sync.WaitGroup

// channel

func goroutineChan(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func channel() {
	// ここに様々な情報を入れて送信して処理する
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int)
	go goroutineChan(s, c)
	x := <-c
	fmt.Println(x)
}

// channel

// Buffer Channels

func bufferChannel() {
	ch := make(chan int, 2)
	ch <- 100
	fmt.Println(len(ch))
	ch <- 200
	fmt.Println(len(ch))
	close(ch)

	for c := range ch {
		fmt.Println(c)
	}
}

// Buffer Channels

// channel range close

func goroutineClose(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		c <- sum
	}
	// ここにclose入れないとエラーになる
	close(c)
}

func channelRange() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int, len(s))
	go goroutineClose(s, c)
	for i := range c {
		fmt.Println(i)
	}
}

// channel range close

// producer consumer

func producer(ch chan int, i int) {
	ch <- i * 2
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		func() {
			defer wg.Done()
			fmt.Println("process", i*1000)
		}()
	}
}

func prpducerConsumer() {
	var wg sync.WaitGroup
	ch := make(chan int)

	// producer
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go producer(ch, i)
	}

	// consumer
	go consumer(ch, &wg)
	wg.Wait()
	close(ch)
}

// producer consumer

// fan out fan in

func multi(first chan int) {
	defer close(first)
	for i := 0; i < 10; i++ {
		first <- i
	}
}

func multi2(first <-chan int, second chan<- int) {
	defer close(second)
	for i := range first {
		second <- i * 2
	}
}

func multi4(second <-chan int, third chan<- int) {
	defer close(third)
	for i := range second {
		third <- i * 4
	}
}

func fanOut() {
	first := make(chan int)
	second := make(chan int)
	third := make(chan int)

	go multi(first)
	go multi2(first, second)
	go multi4(second, third)
	for result := range third {
		fmt.Println(result)
	}
}

// fan out fan in

// select

func goroutineSelect1(ch chan string) {
	for {
		ch <- "packt room 1"
		time.Sleep(1 * time.Second)
	}
}

func goroutineSelect2(ch chan string) {
	for {
		ch <- "packt room 2"
		time.Sleep(1 * time.Second)
	}
}

func selected() {
	c1 := make(chan string)
	c2 := make(chan string)

	go goroutineSelect1(c1)
	go goroutineSelect2(c2)

	for {
		select {
		case message1 := <-c1:
			fmt.Println(message1)
		case message2 := <-c2:
			fmt.Println(message2)
		}
	}
}

// select

// Default Selection for break

func forbreak() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.Tick(500 * time.Millisecond)
	// 名前を指定してbreakしたいところにさすとbreakできる
outLoop:
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			break outLoop
		default:
			fmt.Println("   .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

// Default Selection for break

// sync.Mutex

type Counter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *Counter) Inc(key string) {
	c.mux.Lock()
	// このunLockによって重なりのエラーを回避できる
	defer c.mux.Unlock()
	c.v[key]++
}

func (c *Counter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func syncMutex() {
	c := Counter{v: make(map[string]int)}
	// この書き方で稀に被って実行してしまってエラーが起きる
	go func() {
		for i := 0; i < 10; i++ {
			c.Inc("Key")
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			c.Inc("Key")
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Println(c, c.Value("Key"))
}

// sync.Mutex

// practice

func goroutinePractice(s []string, c chan string) {
	sum := ""
	for _, v := range s {
		sum += v
		c <- sum
	}
	close(c)
}

func practice() {
	words := []string{"test1!", "test2!", "test3!", "test4!"}
	c := make(chan string)
	go goroutinePractice(words, c)
	for w := range c {
		fmt.Println(w)
	}
}

// practice

func main() {
	helloworld()
	channel()
	bufferChannel()
	channelRange()
	prpducerConsumer()
	fanOut()
	selected()
	forbreak()
	syncMutex()
	practice()
}
