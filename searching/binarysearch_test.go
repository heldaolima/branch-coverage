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
		// --- Elemento encontrado ---
		{
			name:     "elemento no meio exato",
			input:    []int{1, 3, 5, 7, 9},
			target:   5,
			expected: 2,
		},
		{
			name:     "elemento no início",
			input:    []int{1, 3, 5, 7, 9},
			target:   1,
			expected: 0,
		},
		{
			name:     "elemento no fim",
			input:    []int{1, 3, 5, 7, 9},
			target:   9,
			expected: 4,
		},
		{
			name:     "elemento em slice com um único item (encontrado)",
			input:    []int{42},
			target:   42,
			expected: 0,
		},
		{
			name:     "elemento com índice 0 em slice grande",
			input:    []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20},
			target:   2,
			expected: 0,
		},
		{
			name:     "elemento no último índice em slice grande",
			input:    []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20},
			target:   20,
			expected: 9,
		},
		{
			name:     "slice com dois elementos — busca pelo primeiro",
			input:    []int{3, 7},
			target:   3,
			expected: 0,
		},
		{
			name:     "slice com dois elementos — busca pelo segundo",
			input:    []int{3, 7},
			target:   7,
			expected: 1,
		},
		{
			name:     "números negativos — elemento encontrado",
			input:    []int{-10, -7, -3, 0, 4},
			target:   -7,
			expected: 1,
		},
		{
			name:     "mix de negativos e positivos — elemento encontrado",
			input:    []int{-5, -1, 0, 3, 8, 13},
			target:   0,
			expected: 2,
		},

		// --- Elemento não encontrado ---
		{
			name:     "target menor que todos os elementos",
			input:    []int{5, 10, 15, 20},
			target:   1,
			expected: -1,
		},
		{
			name:     "target maior que todos os elementos",
			input:    []int{5, 10, 15, 20},
			target:   99,
			expected: -1,
		},
		{
			name:     "target entre elementos existentes (gap)",
			input:    []int{1, 3, 5, 7, 9},
			target:   4,
			expected: -1,
		},
		{
			name:     "elemento em slice com um único item (não encontrado)",
			input:    []int{42},
			target:   7,
			expected: -1,
		},
		{
			name:     "slice vazia",
			input:    []int{},
			target:   1,
			expected: -1,
		},
		{
			name:     "números negativos — elemento não encontrado",
			input:    []int{-10, -7, -3, 0, 4},
			target:   -5,
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

func TestBinarySearchReturnIsValidIndex(t *testing.T) {
	// Garante que o índice retornado realmente aponta para o target na slice
	a := []int{1, 4, 7, 10, 13, 16, 19}
	for _, target := range a {
		idx := searching.BinarySearch(a, target)
		if idx < 0 || idx >= len(a) {
			t.Errorf("índice retornado fora dos limites: %d para target %d", idx, target)
			continue
		}
		if a[idx] != target {
			t.Errorf("a[%d] = %d, mas o target era %d", idx, a[idx], target)
		}
	}
}

func TestBinarySearchNilSlice(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("BinarySearch() causou panic com slice nil: %v", r)
		}
	}()
	got := searching.BinarySearch(nil, 1)
	if got != -1 {
		t.Errorf("BinarySearch(nil, 1) = %d, esperado -1", got)
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	a := make([]int, 1_000_000)
	for i := range a {
		a[i] = i * 2 // 0, 2, 4, ..., 1_999_998
	}
	target := 1_456_780 // elemento presente no meio-final

	for b.Loop() {
		searching.BinarySearch(a, target)
	}
}
