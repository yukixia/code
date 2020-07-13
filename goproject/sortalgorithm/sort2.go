package sortalgorithm

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

}

// Quick3Way 三分取样排序
func Quick3Way(list []int) {

}
