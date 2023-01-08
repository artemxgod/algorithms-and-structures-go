package main

import (
	"log"

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
	a, err := a.DeleteNode(5);
	FatalOnErr(err)

	if a != nil {
		a.PrintPreorder()
	}
}


func FatalOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}