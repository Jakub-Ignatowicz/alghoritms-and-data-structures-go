package main

import (
	"fmt"
	"strconv"

	"github.com/Jakub-Ignatowicz/alghoritms-and-data-structures-go/sorting"
)

type sortingFunc func(arr []int) (uint64, error)

func printSortingAlgorithm(sortFunc sortingFunc, algName string) {
	fmt.Println("\n---------------------")
	fmt.Println(algName)
	fmt.Println("---------------------\n")
	array := []int{64, 34, 25, 12, 22, 11, 90}

	fmt.Println("Array before sorting", array)
	fmt.Println()

	iterationCount, err := sortFunc(array)

	iterationCountStr := strconv.FormatUint(iterationCount, 10)

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Printf("Array sorted with %s: %v\n\n", algName, array)
		fmt.Printf("Array sorted in %s iterations\n", iterationCountStr)
	}
}

func main() {
	printSortingAlgorithm(sorting.BasicBubbleSort, "basic bubble sort")

	printSortingAlgorithm(sorting.DecreasingBubbleSort, "decreasing bubble sort")

	printSortingAlgorithm(sorting.DecreasingFlagBubbleSort, "decreasing flag bubble sort")
}
