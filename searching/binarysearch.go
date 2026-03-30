package searching

func BinarySearch(a []int, target int) int {
	low := 0
	high := len(a) - 1

	for low <= high {
		mid := low + (high-low)/2

		if a[mid] == target {
			return mid
		}

		if a[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
