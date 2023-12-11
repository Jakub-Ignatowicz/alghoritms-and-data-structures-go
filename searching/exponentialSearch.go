package searching

import "math"

func ExponentialSearch(arr []int, value int) (uint64, int) {
	var iterations uint64 = 0

	for i := 0; i < len(arr); i++ {
		iterations++

		index := len(arr) - 1

		if power := int(math.Pow(2, float64(i-1))); power <= len(arr)-1 {
			index = power
		}

		if arr[index] == value {
			return iterations, index
		}

		if arr[index] > value {
			inerIterations, answer := binSearch(arr, value, int(math.Pow(2, float64(i-2))), index, iterations)

			if answer == -1 {
				break
			}

			iterations += inerIterations
			return iterations, answer
		}
	}

	return iterations, -1
}
