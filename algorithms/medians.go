package algorithms

import (
	"fmt"
	"sort"
)

// In computer science, the median of medians is an approximate (median) selection algorithm,
// frequently used to supply a good pivot for an exact selection algorithm,
// mainly the quick-select, that selects kth smallest element.
// Here is source code of the Go Program to Median of Medians to find the Kth Smallest element.
func medianOfMedians(sliceList []int, k, r int) int {
	num := len(sliceList)

	if num < 10 {
		sort.Ints(sliceList)
		return sliceList[k-1]
	}

	// creating slice of medians
	medLen := (num + r - 1) / r
	medians := make([]int, medLen)

	// filling medians slice
	for i := 0; i < medLen; i++ {
		v := i*r + r
		var arr []int
		if v >= num {
			arr = make([]int, len(sliceList[(i*r):]))
			copy(arr, sliceList[(i*r):])
		} else {
			arr = make([]int, r)
			copy(arr, sliceList[(i*r):v])
		}
		medians[i] = arr[len(arr)/2]
	}
	// setting a pivot
	pivot := medianOfMedians(medians, (len(medians)+1)/2, r)

	var left, right []int

	// filling left and right slices relative to pivot
	for i := range sliceList {
		if sliceList[i] < pivot {
			left = append(left, sliceList[i])
		} else if sliceList[i] > pivot {
			right = append(right, sliceList[i])
		}
	}

	switch {
	case k == len(left) + 1:
		return pivot
	case k <= len(left):
		return medianOfMedians(left, k, r)
	default:
		return medianOfMedians(right, k - len(left)-1, r)
	}
}

func TestMedians() {
	intSlice := []int{5, 9, 77, 62, 71, 11, 22, 46, 36, 18, 19, 33, 75, 17, 39, 41, 73, 50, 217, 79, 120}
	// sort to check the algorithm
	sort.Ints(intSlice)
	for _, j := range []int{5, 10, 15, 20} {
		i := medianOfMedians(intSlice, j, 5)
		fmt.Println(j, "th smallest element = ", i)
		v := intSlice[j-1]
		fmt.Println("arr[", j-1, "] = ", v)
		if i != v {
			fmt.Println("Oops! Algorithm is wrong")
		}
	}
}