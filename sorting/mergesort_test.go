package sorting_test

import (
	"branch-cover/sorting"
	"reflect"
	"testing"
)

type SortingTest struct{
	name     string
	input    []int
	expected []int
}

var tests []SortingTest =	[]SortingTest{{
			name:     "slice já ordenada",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "slice em ordem decrescente",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "slice com elementos aleatórios",
			input:    []int{38, 27, 43, 3, 9, 82, 10},
			expected: []int{3, 9, 10, 27, 38, 43, 82},
		},
		{
			name:     "slice com elementos duplicados",
			input:    []int{4, 2, 4, 1, 2},
			expected: []int{1, 2, 2, 4, 4},
		},
		{
			name:     "slice com um único elemento",
			input:    []int{42},
			expected: []int{42},
		},
		{
			name:     "slice com dois elementos fora de ordem",
			input:    []int{9, 1},
			expected: []int{1, 9},
		},
		{
			name:     "slice com dois elementos já ordenados",
			input:    []int{1, 9},
			expected: []int{1, 9},
		},
		{
			name:     "slice com números negativos",
			input:    []int{-3, -1, -7, -4, -2},
			expected: []int{-7, -4, -3, -2, -1},
		},
		{
			name:     "slice mista com negativos e positivos",
			input:    []int{3, -2, 0, -5, 7, 1},
			expected: []int{-5, -2, 0, 1, 3, 7},
		},
		{
			name:     "slice com todos os elementos iguais",
			input:    []int{7, 7, 7, 7},
			expected: []int{7, 7, 7, 7},
		},
		{
			name:     "slice vazia",
			input:    []int{},
			expected: []int{},
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