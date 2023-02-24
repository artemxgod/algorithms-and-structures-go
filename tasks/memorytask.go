package tasks

import "fmt"

func MemoryOne() {
	a := make([]int, 0, 2)
	b := a

	a = append(a, 1)
	b = append(b, 2)
	fmt.Printf("%v %v %p %p", a, b, a, b)

}