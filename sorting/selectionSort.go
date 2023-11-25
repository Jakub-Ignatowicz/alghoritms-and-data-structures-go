package sorting

import "errors"

func SelectionSort(arr []int) (uint64, error) {
	var iterations uint64 = 0
	len := len(arr)
	if len == 0 {
		return 0, errors.New("Array is empty")
	}
	for i := 0; i < len-1; i++ {
		var lowest int
		var lowestIndex int
		for j := i; j < len; j++ {
			iterations++
			if j == i {
				lowest = arr[j]
				lowestIndex = j
			} else if arr[j] < lowest {
				lowest, arr[j] = arr[j], lowest
				lowestIndex, j = j, lowestIndex
			}
		}
		arr[i] = lowest
	}
	return iterations, nil
}
