package sorting

func SelectionSort(arr []int) uint64 {
	n := len(arr)
	var iterations uint64 = 0

	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
			iterations++
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}

	return iterations
}
