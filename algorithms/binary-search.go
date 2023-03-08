package algorithms

func BinarySearch(arr []int, find int) bool {
	low, high := 0, len(arr) - 1

	for low <= high {
		mid := low + (high - low) / 2

		if arr[mid] < find {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if low == len(arr) || arr[low] != find {
		return false
	}

	return true
}