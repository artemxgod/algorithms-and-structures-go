package algorithms

import "fmt"

// This technique pass over the list of elements,
// by using the index to move from the beginning of the list to the end.
// Each element is examined and if it does not match the search item, the next item is examined.
// By hopping from one item to its next, the list is passed over sequentially.
func LinearSearch(arr []int, key int) bool {
	for _, elem := range arr {
		if elem == key {
			return true
		}
	}

	return false
}

// A binary search is a search strategy used to find elements within a list
// by consistently reducing the amount of data to be searched and
// thereby increasing the rate at which the search term is found.
// To use a binary search algorithm, the list to be operated on must have already been sorted.
func BinarySearch(arr []int, find int) bool {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + (high-low)/2

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

// The Interpolation Search is an improvement over Binary Search for instances,
// where the values in a sorted array are uniformly distributed.
// Binary Search always goes to middle element to check.
// On the other hand interpolation search may go to different locations
// according the value of key being searched.
// Here is the source code of the Go program to search element in an integer array
// using Interpolation search algorithm. The output shows the position of element in array.
// Step1: In a loop, calculate the value of “pos” using the probe position formula.
// Step2: If it is a match, return the index of the item, and exit.
// Step3: If the item is less than arr[pos], calculate the probe position of the left sub-array.
func InterpolationSearch(arr []int, key int) int {
	return interpolationReq(arr, key, 0, len(arr)-1)
}

func interpolationReq(arr []int, key, low, high int) int {
	// Since array is sorted, an element present
	// in array must be in range defined by corner
	if low <= high && key >= arr[low] && key <= arr[high] {
		// Probing the position with keeping
		// uniform distribution in mind.
		pos := low + int(float64(high-low)*(float64(key-arr[low])/float64(arr[high]-arr[low])))
		fmt.Println(pos, low, high, arr[pos])

		// Condition of target found
		if arr[pos] == key {
			return pos
		}

		if arr[pos] < key {
			return interpolationReq(arr, key, pos+1, high)
		}

		if arr[pos] > key {
			return interpolationReq(arr, key, low, pos-1)
		}
	}

	return -1
}
