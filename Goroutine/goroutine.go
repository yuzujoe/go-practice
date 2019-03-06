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

func main() {
	// helloworld()
	// channel()
	// bufferChannel()
	channelRange()
}
