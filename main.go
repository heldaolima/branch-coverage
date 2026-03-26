package main

import (
	"branch-cover/sorting"
	"branch-cover/searching"
	"fmt"
)

func main() {
	a := []int{6, 5, 4, 3, 2, 1}
	sorting.MergeSort(a)
	fmt.Printf("%v\n", a)

	b := []int{6, 5, 4, 3, 2, 1}
	sorting.QuickSort(b)
	fmt.Printf("%v\n", b)

	idx := searching.BinarySearch(b, 3)
	fmt.Printf("%v\n", idx)
}