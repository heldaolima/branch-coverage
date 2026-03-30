package sorting_test

import (
	"branch-cover/sorting"
	"reflect"
	"testing"
)

type SortingTest struct {
	name     string
	input    []int
	expected []int
}

var tests []SortingTest = []SortingTest{{
		name:     "slice em ordem decrescente",
		input:    []int{5, 4, 3, 2, 1},
		expected: []int{1, 2, 3, 4, 5},
	},
	{
		name: "slice em ordem aleatória",
		input: []int{65, 3, 634, 23, 50, 10, 0, 2},
		expected: []int{0, 2, 3, 10, 23, 50, 65, 634},
	},
}

func TestMergeSort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
