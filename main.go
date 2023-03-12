package main

import (
	"log"

	"github.com/artemxgod/algorithms-and-structures/algorithms"
)

func main() {
	arr := []int{1, 5, 6, 8, 12, 16, 35, 67, 68, 91, 895}
	log.Println(algorithms.InterpolationSearch(arr, 67))
}

func FatalOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
