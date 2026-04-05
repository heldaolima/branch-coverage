package sorting_test

import (
	"branch-cover/sorting"
	"fmt"
	"reflect"
	"testing"
)

var quickSortTests []SortingTest = []SortingTest{{
	input:    []int{5, 4, 3, 2, 1},
	expected: []int{1, 2, 3, 4, 5},
},
	{
		input:    []int{65, 3, 634, 23, 50, 10, 0, 2},
		expected: []int{0, 2, 3, 10, 23, 50, 65, 634},
	},
	{
		input:    []int{1, 2, 3, 4, 5, 6, 7},
		expected: []int{1, 2, 3, 4, 5, 6, 7},
	},
	{
		input:    []int{3, 3, 3, 3, 3, 3, 3, 3},
		expected: []int{3, 3, 3, 3, 3, 3, 3, 3},
	},
	{
		input: []int{52},
		expected: []int{52},
	},
	{
		input: []int{},
		expected: []int{},
	},
}

func TestQuickSort(t *testing.T) {
	for i, tt := range quickSortTests {
		t.Run(fmt.Sprintf("Teste %d", i), func(t *testing.T) {
			sorting.QuickSort(tt.input)
			if !reflect.DeepEqual(tt.input, tt.expected) {
				t.Errorf("MergeSort() = %v, esperado %v", tt.input, tt.expected)
			}
		})
	}
}

func BenchmarkQuickSort(b *testing.B) {
	base := []int{38, 27, 43, 3, 9, 82, 10, 55, 17, 64, 91, 4, 33, 76, 21}
	for b.Loop() {
		input := make([]int, len(base))
		copy(input, base)
		sorting.QuickSort(input)
	}
}
