package sorting_test

import (
	"branch-cover/sorting"
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
