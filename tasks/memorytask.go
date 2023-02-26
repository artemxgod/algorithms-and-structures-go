package tasks

import (
	"fmt"
)

func MemoryOne() {
	a := make([]int, 0, 2)
	b := a

	a = append(a, 1)
	b = append(b, 2)
	fmt.Printf("%v %v %p %p", a, b, a, b)
}

// Result is a[0] = 2, b[0] = 2 because both slices points on one array but after we append
// the first one, the second one will still have len = 0, so after we append the second slice
// it will put value in b[0].
