package sortalgorithm

import (
	"fmt"
)

// Sort2 归并排序
func Sort2(intervals [][]int) {
	N := len(intervals)
	aux := make([][]int, N, N)
	for sz := 1; sz < N; sz = sz + sz {
		for lo := 0; lo < N-sz; lo = lo + sz + sz {
			high := lo + sz + sz - 1
			if lo+sz+sz > N {
				high = N - 1
			}
			mergeSort(aux, intervals, lo, lo+sz-1, high)
			//fmt.Println(intervals)
		}
	}
}

// UBSort 自顶向下
func UBSort(intervals [][]int) {
	N := len(intervals)
	aux := make([][]int, N, N)
	sort3(aux, intervals, 0, N-1)
}

// Sort3 自顶向下
func sort3(aux [][]int, intervals [][]int, low int, high int) {
	if high <= low {
		return
	}
	mid := (low + high) / 2
	sort3(aux, intervals, low, mid)
	sort3(aux, intervals, mid+1, high)
	mergeSort(aux, intervals, low, mid, high)
}

func mergeSort(aux [][]int, intervals [][]int, low int, mid int, high int) {
	i := low
	j := mid + 1
	for m := low; m <= high; m++ {
		aux[m] = intervals[m]
	}
	//fmt.Println(aux, low, mid, high)
	for k := low; k <= high; k++ {
		if i > mid {
			intervals[k] = aux[j]
			j = j + 1
		} else if j > high {
			intervals[k] = aux[i]
			i = i + 1
		} else if aux[j][0] < aux[i][0] {
			intervals[k] = aux[j]
			j = j + 1
		} else {
			intervals[k] = aux[i]
			i = i + 1
		}
	}
}

// QuickSort 快速排序
func QuickSort(list []int) {
	quicksort(list, 0, len(list)-1)

}

func quicksort(list []int, low int, high int) {
	if low >= high {
		return
	}
	// i := porition(list, low, high)
	// fmt.Printf("%+v\n", list)
	// quicksort(list, low, i-1)
	// quicksort(list, i+1, high)
	lt, gt := quick3way(list, low, high)
	quicksort(list, low, lt-1)
	quicksort(list, gt+1, high)
}

func porition(list []int, low int, high int) int {
	pivot := list[low]
	i := low
	j := high
	for true {
		for list[i] < pivot {
			i++
		}
		for list[j] > pivot {
			j--
		}
		if i >= j {
			break
		}
		temp := list[j]
		list[j] = list[i]
		list[i] = temp
		//list[i], list[j] = list[j], list[i]

	}
	//list[low], list[j] = list[j], list[low]
	temp := list[j]
	list[j] = list[low]
	list[i] = temp
	return j
}

// Quick3Way 三分取样排序
func quick3way(list []int, low int, high int) (int, int) {
	lt := low
	gt := high
	pivot := list[low]
	i := low + 1
	for i <= gt {
		if list[i] < pivot {
			temp := list[i]
			list[i] = list[lt]
			list[lt] = temp
			lt++
			i++
		} else if list[i] > pivot {
			temp := list[i]
			list[i] = list[gt]
			list[gt] = temp
			gt--
		} else {
			i++
		}
	}
	return lt, gt
}

func getMax(list []int) int {
	max := list[0]
	for i := 1; i < len(list); i++ {
		if max < list[i] {
			max = list[i]
		}
	}
	return max
}

/*
	桶排序
	适用于数据范围n不是很大的数量排序
	思路：将数据分成k=n/m个桶，每个桶的数据量为n/m,在对每个桶使用快速排序即可
*/
func BucketSort(list []int) {
	num := len(list)
	if num <= 1 {
		return
	}
	max := getMax(list)
	bucketNum := max
	buckets := make([][]int, bucketNum) //二维切片
	//分桶
	for i := 0; i < num; i++ {
		index := list[i] * (bucketNum - 1) / max
		buckets[index] = append(buckets[index], list[i])
	}
	tmppos := 0
	//对每个桶内的数据进行排序
	for i := 0; i < bucketNum; i++ {
		fmt.Printf("%d\n", len(buckets[i]))
		if len(buckets[i]) > 1 {
			QuickSort(buckets[i])
		}
		copy(list[tmppos:], buckets[i])
		tmppos += len(buckets[i])
	}
}

// func TestSort2(t *testing.T) {
// 	list := []int{1, 6, 3, 5, 8, 6, 4, 5, 12, 4, 7, 8, 13, 2, 4}
// 	fmt.Printf("%+v\n", list)
// 	BucketSort(list)
// 	fmt.Printf("%+v\n", list)
// }
