package main

import (
	"log"

	"github.com/artemxgod/algorithms-and-structures/structures"
)

func main() {
	list := structures.List{}

	list.Insert(4)
	list.Insert(5)
	list.Insert(123)
	list.Insert("booga")
	structures.PrintList(list.Head)
	structures.PrintListBackwords(list.Tail)
	list.Reverse()
	list.Display()
	list.Push("Amen")
	list.Display()
	structures.PrintList(list.Head)
	structures.PrintListBackwords(list.Tail)
	
}

func FatalOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
