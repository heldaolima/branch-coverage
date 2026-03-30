package searching_test

import (
	"branch-cover/searching"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		target   int
		expected int
	}{
		{
			name:     "elemento no meio exato",
			input:    []int{1, 3, 5, 7, 9},
			target:   5,
			expected: 2,
		},
		{
			name: "elemento na esquerda",
			input:    []int{1, 3, 5, 7, 9},
			target: 3,
			expected: 1,
		},
		{
			name: "elemento não encontrado",
			input:    []int{1, 3, 5, 7, 9},
			target: 20,
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := searching.BinarySearch(tt.input, tt.target)
			if got != tt.expected {
				t.Errorf("BinarySearch(%v, %d) = %d, esperado %d", tt.input, tt.target, got, tt.expected)
			}
		})
	}
}