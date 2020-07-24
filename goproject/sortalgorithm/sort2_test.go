package sortalgorithm

import (
	"fmt"
	"testing"
)

func TestSort2(t *testing.T) {
	list := []int{4, 3, 6, 5, 7, 1, 8, 9}
	fmt.Printf("%+v", list)
	QuickSort(list)
	fmt.Printf("%+v", list)
}
