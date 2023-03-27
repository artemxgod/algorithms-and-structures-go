package main

import (
	"log"

	"github.com/artemxgod/algorithms-and-structures/toolbox/cmd"
)

func main() {
	// structures.TestLinkedList()
	cmd.Execute()
}

func FatalOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// P.S. We can add mutex into struct
// IF MUTEX IS LOCKED IT MUST BE UNLOCKED
