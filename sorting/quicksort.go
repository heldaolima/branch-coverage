package sorting

import "math/rand"

func QuickSort(a []int) {
	quickSort(a, 0, len(a)-1)
}

func quickSort(a []int, low, high int) {
	if low < high {
		p := partition(a, low, high)

		quickSort(a, low, p-1)
		quickSort(a, p+1, high)
	}
}

func partition(a []int, low, high int) int {
	pivotIndex := rand.Intn(high-low+1) + low // pivot is random in [low, high]

	a[pivotIndex], a[high] = a[high], a[pivotIndex]

	pivot := a[high]
	i := low

	for j := low; j < high; j++ {
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}

	a[i], a[high] = a[high], a[i]
	return i
}
