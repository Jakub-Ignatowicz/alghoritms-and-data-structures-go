package searching

func LinearSearch(arr []int, value int) (uint64, int) {
	var iterations uint64 = 0

	i := 0

	for ; i < len(arr); i++ {
		iterations++
		if arr[i] == value {
			return iterations, i
		}
	}

	return iterations, -1
}
