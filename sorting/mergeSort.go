package sorting

type sortingFunc func([]int) uint64

func MergeSort(arr []int) uint64 {
	if len(arr) <= 1 {
		return 0
	}

	mid := len(arr) / 2
	left := make([]int, mid)
	right := make([]int, len(arr)-mid)

	copy(left, arr[:mid])
	copy(right, arr[mid:])

	var iterations uint64

	iterations += MergeSort(left)
	iterations += MergeSort(right)
	iterations += merge(arr, left, right)

	return iterations
}

func merge(arr, left, right []int) uint64 {
	var i, j, k uint64
	iterations := uint64(0)

	for i < uint64(len(left)) && j < uint64(len(right)) {
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
		iterations++
	}

	for i < uint64(len(left)) {
		arr[k] = left[i]
		i++
		k++
	}

	for j < uint64(len(right)) {
		arr[k] = right[j]
		j++
		k++
	}

	return iterations
}
