package main

import (
	"fmt"
	"log"

	"github.com/artemxgod/algorithms-and-structures/packages"
)

func main() {
	// structures.TestLinkedList()
	// cmd.Execute()
	// structures.BitOperators()
	// packages.Poll()
	procsID := packages.ListProcesses()
	fmt.Println(procsID)

	names := packages.ProcsNameByID(procsID)
	fmt.Println("NAMES:", names[0])
}

func FatalOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
