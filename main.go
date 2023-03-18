package main

import (
	"log"

	"github.com/artemxgod/algorithms-and-structures/concurrency"
)

func main() {
	// structures.TestLinkedList()
	concurrency.TestSelection()
}

func FatalOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
