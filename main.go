package main

import (
	"fmt"
	"log"

	"github.com/artemxgod/algorithms-and-structures/algorithms"
)

func main() {
	// arr := []int{1, 5, 6, 8, 12, 16, 35, 67, 68, 91, 895}
	// log.Println(algorithms.InterpolationSearch(arr, 67))
	txt := "Australia is a country and continent surrounded by the Indian and Pacific oceans."
	patterns := []string{"and", "the", "surround", "Pacific", "Germany", "Australia"}
	fmt.Println(algorithms.RKSearch(txt, patterns))
}

func FatalOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
