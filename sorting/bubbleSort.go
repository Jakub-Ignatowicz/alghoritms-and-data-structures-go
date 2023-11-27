package sorting

func BasicBubbleSort(arr []int) uint64 {
	n := len(arr)

	var iterationCount uint64 = 0

	if n <= 1 {
		return 0
	}

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1; j++ {
			iterationCount++

			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return iterationCount
}

func DecreasingBubbleSort(arr []int) uint64 {
	n := len(arr)

	var iterationCount uint64 = 0

	if n <= 1 {
		return 0
	}

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			iterationCount++

			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return iterationCount
}

func DecreasingFlagBubbleSort(arr []int) uint64 {
	n := len(arr)

	var iterationCount uint64 = 0

	if n <= 1 {
		return 0
	}

	for i := 0; i < n-1; i++ {
		changed := false
		for j := 0; j < n-i-1; j++ {
			iterationCount++

			if arr[j] > arr[j+1] {
				changed = true
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		if !changed {
			break
		}
	}
	return iterationCount
}
