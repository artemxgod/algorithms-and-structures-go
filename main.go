package main

import (
	"log"

	"github.com/artemxgod/algorithms-and-structures/structures"
)

func main() {
	// structures.TestLinkedList()
	// cmd.Execute()
	structures.BitOperators()
}

func FatalOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
