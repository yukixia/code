package leecode

import "fmt"

func twoSum(nums []int, target int) []int {
	ret := []int{-1, -1}
	dict := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if nums[i] <= target {
			dict[nums[i]] = i
		}
	}
	for j := 0; j < len(nums); j++ {
		secondValue := target - nums[j]
		secondIndex, exists := dict[secondValue]
		if exists {
			fmt.Println(secondValue)
			ret[0] = j
			ret[1] = secondIndex
			break
		}
	}
	return ret
}
