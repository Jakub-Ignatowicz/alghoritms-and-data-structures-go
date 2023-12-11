package searching

func InterpolationSearch(arr []int, value int) (uint64, int) {
	var iterations uint64 = 0

	low, high := 0, len(arr)-1

	for low <= high && arr[low] <= value && arr[high] >= value {
		iterations++

		pos := low + ((value - arr[low]) * (high - low) / (arr[high] - arr[low]))

		if arr[pos] == value {
			return iterations, pos
		}

		if arr[pos] < value {
			low = pos + 1
		} else {
			high = pos - 1
		}
	}

	return iterations, -1
}
