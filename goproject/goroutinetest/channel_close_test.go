package goroutinetest

import (
	"fmt"
	"sync"
	"testing"
)

func dataProducer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		if data, ok := <-ch; ok {
			fmt.Println(data)
		} else {
			fmt.Println("channel close.")
			break
		}
	}
}

func TestChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	go dataProducer(ch, &wg)
	wg.Add(1)
	go dataReceiver(ch, &wg)
	wg.Add(1)
	go dataReceiver(ch, &wg)
	wg.Wait()
}
