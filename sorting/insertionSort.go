package sorting

func InsertionSort(arr []int) uint64 {
	length := len(arr)
	var iterations uint64 = 0
	if length <= 1 {
		return iterations
	}

	for i := 1; i < length; i++ {
		for j := i - 1; j >= 0 && arr[j+1] < arr[j]; j-- {
			arr[j+1], arr[j] = arr[j], arr[j+1]
			iterations++
		}
	}

	return iterations
}
