package main

import (
	"log"

	"github.com/artemxgod/algorithms-and-structures/algorithms"
)

func main() {
	// structures.TestLinkedList()
	algorithms.TestKMPSearch()
}

func FatalOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
