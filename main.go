package main

import (

	"github.com/artemxgod/algorithms-and-structures/structures"
)

func main() {
	a := structures.NewTree(5)
	a.Insert(6)
	a.Insert(9)
	a.Insert(3)
	a.Insert(1)
	a.Insert(4)

	
	a.PrintTree()
}