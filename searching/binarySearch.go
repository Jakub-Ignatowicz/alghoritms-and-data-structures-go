package searching

func BinarySearch(arr []int, value int) (uint64, int) {
	return binSearch(arr, value, 0, len(arr)-1, 0)
}

func binSearch(arr []int, value int, begin int, end int, iterations uint64) (uint64, int) {
	midIndex := (begin + end) / 2
	iterations++

	if arr[midIndex] == value {
		return iterations, midIndex
	}

	if begin != end {
		if arr[midIndex] > value {
			return binSearch(arr, value, begin, midIndex-1, iterations)
		} else {
			return binSearch(arr, value, midIndex+1, end, iterations)
		}
	}

	return iterations, -1
}
