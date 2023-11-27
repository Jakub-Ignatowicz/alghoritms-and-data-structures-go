package sorting

import "errors"

func CountingSort(arr []int) (uint64, error) {
	if len(arr) == 0 {
		return 0, errors.New("Array is empty")
	}
	var iterationCount uint64 = 0
	var aMin, aMax = arr[0], arr[0]
	for _, x := range arr {
		if x < aMin {
			aMin = x
		}
		if x > aMax {
			aMax = x
		}
		iterationCount++
	}
	count := make([]int, aMax-aMin+1)
	for _, x := range arr {
		count[x-aMin]++
	}
	z := 0
	for i, c := range count {
		for c > 0 {
			arr[z] = i + aMin
			z++
			c--
			iterationCount++
		}
	}
	return iterationCount, nil
}
