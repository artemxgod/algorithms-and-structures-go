package main

import (
	"log"

	"github.com/artemxgod/algorithms-and-structures/packages"
)

func main() {
	// structures.TestLinkedList()
	packages.TestProtoBuf()
}

func FatalOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
