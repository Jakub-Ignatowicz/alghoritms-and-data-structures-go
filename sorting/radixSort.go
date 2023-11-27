package sorting

func radixCountingSortVariation(arr []int, exp int) uint64 {
	var digits [10]int
	var iterations uint64 = 0
	lenght := len(arr)
	var ans = make([]int, lenght)

	var digitsPosition [10]int
	for _, val := range arr {
		digits[(val/exp)%10]++
		iterations++
	}

	currVal := 0
	for index := range digitsPosition {
		currVal += digits[index]
		digitsPosition[index] = currVal
		iterations++
	}

	for i := lenght - 1; i >= 0; i-- {
		digitValue := (arr[i] / exp) % 10

		digitsPosition[digitValue]--
		ans[digitsPosition[digitValue]] = arr[i]
		iterations++
	}
	for index := range ans {
		arr[index] = ans[index]
	}
	return iterations
}

func RadixSort(arr []int) uint64 {
	var finalIterations uint64 = 0
	maxVal := arr[0]

	for _, num := range arr {
		if maxVal < num {
			maxVal = num
		}
	}

	for exp := 1; maxVal/exp > 0; exp *= 10 {
		var iterations uint64
		iterations = radixCountingSortVariation(arr, exp)
		finalIterations += iterations
	}

	return finalIterations
}
