package algorithms

// Time Complexity - O(n^2)
// It swaps two neighboring elements from arr[0] to arr[n-1] so every inner loop the max value [0:n-idx] will stand on its place
// thats why we decrease max jdx by 1 every time
// We stop if inner array does not swap any elements
func BubbleSort(arr []int) {
	n := len(arr)

	sorted := false
	for idx := 0; !sorted; idx++ {
		sorted = true
		for jdx := 0; jdx < n-1-idx; jdx++ {
			if arr[jdx] > arr[jdx+1] {
				arr[jdx], arr[jdx+1] = arr[jdx+1], arr[jdx]
				sorted = false
			}
		}
	}
}

// Time Complexity - O(n^2)
// Selection sort find min element [idx+1:n] and puts it in arr[idx]
func SelectionSort(arr []int) {
	n := len(arr)

	for idx := 0; idx < n - 1; idx++ {
		min_idx := idx
		for jdx := idx+1; jdx < n; jdx++ {
			if arr[jdx] < arr[min_idx] {
				min_idx = jdx
			}
		}
		if min_idx != idx {
			arr[min_idx], arr[idx] = arr[idx], arr[min_idx]
		}
	}
}

// Time Complexity - O(n^2)
// Sorts two element, than every iteration takes one more unsorted element and compare it with sorted part
func InsertionSort(arr []int) {
	n := len(arr)

	for idx := 1; idx < n; idx++ {
		for jdx := idx; jdx > 0 && arr[jdx-1] > arr[jdx]; jdx-- {
			arr[jdx], arr[jdx-1] = arr[jdx-1], arr[jdx]
		}
	} 
}

// Time Complexity - O(n^2)
// Merge sort divide array into smaller arrays and sorts them
func MergeSort(arr []int) {
	mergeReq(arr, 0, len(arr)-1) 
}

func mergeReq(arr []int, start, end int) {
	if start < end {
		// same as (start+end)/2 but avoids overflow for large numbers
		mid := start + (end - start) / 2

		// separating 
		mergeReq(arr, start, mid)
		mergeReq(arr, mid + 1, end)

		// merge separated arrays
		merge(arr, start, mid, end)
	}
}

func merge(arr []int, start, mid, end int) {
	lenLeft, lenRight := mid - start + 1, end - mid

	// creating sub arrays
	left := make([]int, lenLeft)
	right := make([]int, lenRight)

	// copy array parts into subarrays
	copy(left, arr[start:mid+1])
	copy(right, arr[mid+1:])
	
	var i, j, k int
	k = start

	// merge subarrays back into arr
	for i < lenLeft && j < lenRight {
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}

	// copy remaining elements into main array
	for i < lenLeft {
		arr[k] = left[i]
		i++; k++;
	}
	for j < lenRight {
		arr[k] = right[j]
		j++; k++;
	}
}

