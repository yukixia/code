// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var wg sync.WaitGroup

// func main() {
// 	wg.Add(2)
// 	fmt.Println("Start Goroutines")
// 	go printPrime("A")
// 	go printPrime("B")
// 	fmt.Println("Waitting Goroutine")
// 	wg.Wait()
// 	fmt.Println("Terminating program")
// }

// func printPrime(prefix string) {
// 	defer wg.Done()
// next:
// 	for outer := 2; outer <= 5000; outer++ {
// 		for inner := 2; inner <= outer; inner++ {
// 			if outer%inner == 0 {
// 				continue next
// 			}
// 		}
// 		fmt.Printf("%s:%d\n", prefix, outer)
// 	}
// 	fmt.Println("Complete", prefix)
// }