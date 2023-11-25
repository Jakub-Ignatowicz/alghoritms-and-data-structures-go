package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Jakub-Ignatowicz/alghoritms-and-data-structures-go/sorting"
)

type sortingFunc func(arr []int) (uint64, error)

func printSortingAlgorithm(sortFunc sortingFunc, algName string) {
	fmt.Println("\n---------------------")
	fmt.Println(algName)
	fmt.Println("---------------------\n")
	array := []int{56, 23, 78, 12, 45, 67, 89, 34, 90, 14, 76, 29, 50, 88, 41, 3, 19, 72, 5, 61, 98, 27, 83, 9, 38, 71, 2, 64, 49, 94}

	fmt.Println("Array before sorting", array)
	fmt.Println()

	start := time.Now()
	iterationCount, err := sortFunc(array)
	end := time.Now()
	duration := end.Sub(start)

	iterationCountStr := strconv.FormatUint(iterationCount, 10)

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Printf("Array sorted with %s: %v\n\n", algName, array)
		fmt.Printf("Array sorted in %s iterations\n", iterationCountStr)
		fmt.Printf("Arrat sorted in %.2f nanoseconds", float64(duration.Nanoseconds()))
	}
}

func main() {
	printSortingAlgorithm(sorting.BasicBubbleSort, "basic bubble sort")

	printSortingAlgorithm(sorting.DecreasingBubbleSort, "decreasing bubble sort")

	printSortingAlgorithm(sorting.DecreasingFlagBubbleSort, "decreasing flag bubble sort")

	printSortingAlgorithm(sorting.SelectionSort, "selection sort")
}
