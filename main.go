package main

import (
	"log"

	"github.com/artemxgod/algorithms-and-structures/tasks"
)

func main() {
	tasks.Acid()
}

func FatalOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
