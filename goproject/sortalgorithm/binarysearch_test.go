package sortalgorithm

import (
	"fmt"
	"testing"
)

func BinarySearch(list []int, target int) int {
	QuickSort(list)
	fmt.Printf("%+v\n", list)
	low := 0
	high := len(list) - 1
	for low <= high {
		mid := (low + high) / 2
		if target < list[mid] {
			high = mid - 1
		} else if target > list[mid] {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func TestBinarySearch(t *testing.T) {
	list := []int{4, 2, 7, 5, 8, 1, 9, 10, 6}
	index := BinarySearch(list, 10)
	t.Log(index)
}
