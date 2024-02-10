package main

import (
	"log"

	"github.com/artemxgod/algorithms-and-structures/packages"
)

func main() {
	// structures.TestLinkedList()
	// cmd.Execute()
	// structures.BitOperators()
	packages.Poll()
}

func FatalOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
