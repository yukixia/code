package goroutinetest

import (
	"fmt"
	"sync"
	"testing"
)

type Singleton struct {
}

var singletonInstance *Singleton
var once sync.Once

func GetSingleObj() *Singleton {
	once.Do(func() {
		fmt.Println("Create obj.")
		singletonInstance = new(Singleton)
	})
	return singletonInstance
}

func TestGetSingleObj(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingleObj()
			fmt.Printf("%p\n", obj)
			wg.Done()
		}()
	}
	wg.Wait()
}
