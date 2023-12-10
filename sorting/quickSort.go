package sorting

import "sync"

func QuickSortLomuto(arr []int) uint64 {
	return lomutoPartition(arr, 0, len(arr)-1)
}

func QuickSortLomutoMultiThreaded(arr []int) uint64 {
	return lomutoThread(arr, 0, len(arr)-1)
}

func lomutoPartition(arr []int, low int, high int) uint64 {
	var iterations uint64 = 0
	if low < high {
		pivot := arr[high]
		j := low

		for i := low; i < high; i++ {
			if arr[i] <= pivot {
				arr[i], arr[j] = arr[j], arr[i]
				j++
			}
			iterations++
		}

		arr[high], arr[j] = arr[j], arr[high]

		iterations += lomutoPartition(arr, low, j-1)
		iterations += lomutoPartition(arr, j+1, high)
	}

	return iterations
}

func lomutoThread(arr []int, low int, high int) uint64 {
	var iterations uint64 = 0
	if low < high {
		pivot := arr[high]
		j := low

		for i := low; i < high; i++ {
			if arr[i] <= pivot {
				arr[i], arr[j] = arr[j], arr[i]
				j++
			}
			iterations++
		}

		arr[high], arr[j] = arr[j], arr[high]

		var wg sync.WaitGroup

		wg.Add(2)

		go func() {
			defer wg.Done()
			iterations += lomutoPartition(arr, low, j-1)
		}()
		go func() {
			defer wg.Done()
			iterations += lomutoPartition(arr, j+1, high)
		}()

		wg.Wait()
	}

	return iterations
}

func HoareQuickSort(arr []int) uint64 {
	var iterationCount uint64
	iterationCount = 0
	hoareSort(arr, 0, len(arr)-1, &iterationCount)
	return iterationCount
}

func hoareSort(arr []int, low, high int, iterationCount *uint64) {
	if low < high {
		pivotIndex := hoarePartition(arr, low, high, iterationCount)

		hoareSort(arr, low, pivotIndex, iterationCount)
		hoareSort(arr, pivotIndex+1, high, iterationCount)
	}
}

func hoarePartition(arr []int, low, high int, iterationCount *uint64) int {
	pivot := arr[low]
	i := low - 1
	j := high + 1

	for {
		for {
			i++
			*iterationCount++
			if arr[i] >= pivot {
				break
			}
		}

		for {
			j--
			*iterationCount++
			if arr[j] <= pivot {
				break
			}
		}

		if i >= j {
			return j
		}

		arr[i], arr[j] = arr[j], arr[i]
	}
}
