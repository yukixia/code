// package main

// import (
// 	"fmt"
// 	"runtime"
// 	"sync"
// )

// var (
// 	counter int
// 	wg      sync.WaitGroup
// 	mutex   sync.Mutex
// )

// func main() {
// 	wg.Add(2)
// 	go incCounter(1)
// 	go incCounter(2)
// 	wg.Wait()
// 	fmt.Printf("Final Counter:%d \n", counter)
// }

// func incCounter(id int) {
// 	defer wg.Done()
// 	for i := 0; i < 2; i++ {
// 		mutex.Lock()
// 		{
// 			value := counter
// 			runtime.Gosched()
// 			value++
// 			counter = value
// 		}
// 		mutex.Unlock()
// 	}
// }
