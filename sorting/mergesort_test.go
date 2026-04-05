package sorting_test

import (
	"branch-cover/sorting"
	"fmt"
	"reflect"
	"testing"
)

type SortingTest struct {
	input    []int
	expected []int
}

var mergeSortTests []SortingTest = []SortingTest{{
	input:    []int{6, 5, 4, 3, 2, 1},
	expected: []int{1, 2, 3, 4, 5, 6},
},
	{
		input:    []int{65, 3, 634, 23, 50, 10, 0, 2},
		expected: []int{0, 2, 3, 10, 23, 50, 65, 634},
	},
	{
		input:    []int{-4, -6, -10, -8, -5, -1},
		expected: []int{-10, -8, -6, -5, -4, -1},
	},
	{
		input:    []int{3, 2, 1, 4, 5, 6},
		expected: []int{1, 2, 3, 4, 5, 6},
	},
	{
		input:    []int{5, 5, 5, 5, 5, 5},
		expected: []int{5, 5, 5, 5, 5, 5},
	},
	{
		input:    []int{1, 2, 3, 4, 5, 6},
		expected: []int{1, 2, 3, 4, 5, 6},
	},
	{
		input:    []int{-10, -8, 3, 7, 48, -59, 50},
		expected: []int{-59, -10, -8, 3, 7, 48, 50},
	},
	{
		input:    []int{2, 2, 2, 1},
		expected: []int{1, 2, 2, 2},
	},
	{
		input:    []int{9, 7, 5, 3, 1, 8, 6},
		expected: []int{1, 3, 5, 6, 7, 8, 9},
	},
	{
		input: []int{42},
		expected: []int{42},
	},
	{
		input: []int{},
		expected: []int{},
	},
}

func TestMergeSort(t *testing.T) {
	for i, tt := range mergeSortTests {
		t.Run(fmt.Sprintf("Teste %d", i), func(t *testing.T) {
			sorting.MergeSort(tt.input)
			if !reflect.DeepEqual(tt.input, tt.expected) {
				t.Errorf("MergeSort() = %v, esperado %v", tt.input, tt.expected)
			}
		})
	}
}

func BenchmarkMergeSort(b *testing.B) {
	base := []int{38, 27, 43, 3, 9, 82, 10, 55, 17, 64, 91, 4, 33, 76, 21}
	for b.Loop() {
		input := make([]int, len(base))
		copy(input, base)
		sorting.MergeSort(input)
	}
}
