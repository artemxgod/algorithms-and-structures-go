package main

import (
	"log"

	"github.com/artemxgod/algorithms-and-structures/algorithms"
)

func main() {
	arr := []int{5, 2, 4, 7, 1, 7, 5, 10, 9, 22}
	// arr2 := []int{5, 4, 3, 2, 1}
	
	log.Println(arr)
	algorithms.MergeSort(arr)
	log.Println(arr)
}

func FatalOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
