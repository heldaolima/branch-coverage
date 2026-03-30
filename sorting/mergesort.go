package sorting

func MergeSort(a []int) {
	aux := make([]int, len(a))
	copy(aux, a)
	mergeSort(a, aux, 0, len(a)-1)
}

func mergeSort(a, aux []int, begin, end int) {
	if begin >= end {
		return
	}

	mid := begin + (end-begin)/2
	mergeSort(a, aux, begin, mid)
	mergeSort(a, aux, mid+1, end)
	merge(a, aux, begin, mid, end)
}

func merge(a, aux []int, begin, mid, end int) {
	for i := begin; i <= end; i++ {
		aux[i] = a[i]
	}

	i, j := begin, mid+1

	for k := begin; k <= end; k++ {
		if i > mid { // Left side is over
			a[k] = aux[j]
			j++
		} else if j > end { // right side is over
			a[k] = aux[i]
			i++
		} else if aux[i] <= aux[j] { // comparison
			a[k] = aux[i]
			i++
		} else {
			a[k] = aux[j]
			j++
		}
	}
}
