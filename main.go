package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Jakub-Ignatowicz/alghoritms-and-data-structures-go/searching"
	"github.com/Jakub-Ignatowicz/alghoritms-and-data-structures-go/sorting"
)

type sortingFunc func(arr []int) uint64
type searchingFunc func(arr []int, value int) (uint64, int)

func printSearchingAlgorithm(searchFunc searchingFunc, algName string) {
	fmt.Printf("\n\n---------------------\n%s\n---------------------\n\n", algName)

	searchVal := 88

	array := []int{2, 3, 5, 9, 12, 14, 19, 23, 27, 29, 34, 38, 41, 45, 49, 50, 56, 61, 64, 67, 71, 72, 76, 78, 83, 88, 89, 90, 94, 98}

	start := time.Now()
	iterationCount, index := searchFunc(array, searchVal)
	end := time.Now()
	duration := end.Sub(start)

	if index == -1 {
		fmt.Println("Value hasn't been found in this array")
	} else {
		fmt.Printf("Index of value %v was found with %s in %v iterations and %.2f nanoseconds\n", searchVal, algName, iterationCount, float64(duration.Nanoseconds()))
	}
}

func printSortingAlgorithm(sortFunc sortingFunc, algName string) {
	fmt.Printf("\n\n---------------------\n%s\n---------------------\n\n", algName)

	array := []int{56, 23, 78, 12, 45, 67, 89, 34, 90, 14, 76, 29, 50, 88, 41, 3, 19, 72, 5, 61, 98, 27, 83, 9, 38, 71, 2, 64, 49, 94}

	fmt.Println("Array before sorting", array)
	fmt.Println()

	start := time.Now()
	iterationCount := sortFunc(array)
	end := time.Now()
	duration := end.Sub(start)

	iterationCountStr := strconv.FormatUint(iterationCount, 10)

	fmt.Printf("Array sorted with %s: %v\n\n", algName, array)
	fmt.Printf("Array sorted in %s iterations\n", iterationCountStr)
	fmt.Printf("Arrat sorted in %.2f nanoseconds", float64(duration.Nanoseconds()))
}

func main() {
	printSortingAlgorithm(sorting.BasicBubbleSort, "basic bubble sort")

	printSortingAlgorithm(sorting.DecreasingBubbleSort, "decreasing bubble sort")

	printSortingAlgorithm(sorting.DecreasingFlagBubbleSort, "decreasing flag bubble sort")

	printSortingAlgorithm(sorting.SelectionSort, "selection sort")

	printSortingAlgorithm(sorting.CountingSort, "counting sort")

	printSortingAlgorithm(sorting.RadixSort, "radix sort")

	printSortingAlgorithm(sorting.InsertionSort, "insertion sort")

	printSortingAlgorithm(sorting.MergeSort, "merge sort")

	printSortingAlgorithm(sorting.MergeSortMultiThreaded, "multithreaded merge sort")

	printSortingAlgorithm(sorting.QuickSortLomuto, "quick sort lomuto")

	printSortingAlgorithm(sorting.QuickSortLomutoMultiThreaded, "multithreaded quick sort lomuto")

	printSortingAlgorithm(sorting.HoareQuickSort, "quick sort hoare")

	fmt.Println("\n\n------------------------------------------------------------------------------------------------------------------------------")
	fmt.Println()
	fmt.Println("                                                         SEARCHING                                                           ")
	fmt.Println()
	fmt.Println("------------------------------------------------------------------------------------------------------------------------------")

	printSearchingAlgorithm(searching.LinearSearch, "linear search")

	printSearchingAlgorithm(searching.BinarySearch, "binary search")

	printSearchingAlgorithm(searching.InterpolationSearch, "interpolation search")
}
