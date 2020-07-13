package array

import (
	"fmt"
	"testing"
)

type user struct {
	name string
	age  int
}

func testarray(arr *[5]int) {
	arr[0] = 100
	fmt.Println("input testarray,", arr)
}

func testslice(slice []int) {
	slice[0] = 100
	fmt.Println("input testslice,", slice)
}

func testobject(usertest *user) {
	usertest.age = 10
	fmt.Println("input testobject,", *usertest)
}

func testmap(maptest map[string]int) {
	maptest["test2"] = 100
	fmt.Println("input testmap,", maptest)
}

func TestArray(t *testing.T) {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("input testarray before,", arr)
	testarray(&arr)
	fmt.Println("input testarray after,", arr)
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("input testarray before,", slice)
	testslice(slice)
	fmt.Println("input testslice after,", slice)
	usertest := user{
		name: "yunxia",
		age:  23,
	}
	fmt.Println("input testobject before", usertest)
	testobject(&usertest)
	fmt.Println("input testobject after,", usertest)
	maptest := map[string]int{"test1": 1, "test2": 2}
	fmt.Println("input test map_test before,", maptest)
	testmap(maptest)
	fmt.Println("input testmap after,", maptest)
}
