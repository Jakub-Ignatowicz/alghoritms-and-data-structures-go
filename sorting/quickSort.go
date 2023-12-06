package sorting

func QuickSortLomuto(arr []int) uint64 {
	return lomuto(arr, 0, len(arr)-1)
}

func lomuto(arr []int, low int, high int) uint64 {
	var iterations uint64 = 0
	if low < high {
		pivot := arr[high]
		j := low

		for i := low; i < high; i++ {
			if arr[i] <= pivot {
				arr[i], arr[j] = arr[j], arr[i]
				j++
			}
			iterations++
		}

		arr[high], arr[j] = arr[j], arr[high]

		iterations += lomuto(arr, low, j-1)
		iterations += lomuto(arr, j+1, high)
	}

	return iterations
}
