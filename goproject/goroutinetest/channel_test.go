package goroutinetest

import (
	"fmt"
	"testing"
	"time"
)

func Service() string {
	time.Sleep(time.Millisecond * 500)
	return "Done"
}

func OtherService() {
	fmt.Println("otherservice begin")
	time.Sleep(time.Millisecond * 50)
	fmt.Println("otherservice done")
}

func AsyncService() chan string {
	retch := make(chan string)
	go func() {
		ret := Service()
		fmt.Println("return service")
		retch <- ret
		fmt.Println("service exited.")
	}()
	return retch
}

func TestChannel(t *testing.T) {
	select {
	case retch := <-AsyncService():
		fmt.Println(retch)
	case <-time.After(time.Millisecond * 100):
		fmt.Println("time out")
	}
}
