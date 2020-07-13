package sortalgorithm

// SelectSort 选择排序算法
func SelectSort(list []int) {
	listLen := len(list)
	if 0 == listLen {
	}
	for i := 0; i < listLen; i++ {
		min := i
		for j := i + 1; j < listLen; j++ {
			if list[min] > list[j] {
				min = j
			}
		}
		temp := list[i]
		list[i] = list[min]
		list[min] = temp
	}
}
