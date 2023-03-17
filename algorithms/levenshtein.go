package algorithms

import "fmt"

// Levenshtein distance (LD) is a measure of the similarity between two strings,
// which we will refer to as the source string (s) and the target string (t).
// The distance is the number of deletions, insertions,
// or substitutions required to transform s into t.
// For example, the Levenshtein distance between "Asheville" and "Arizona" is 8.
// Here is source code of the Go Program to Implement Levenshtein Distance Computing Algorithm.
func LevenshteinDistance(src, target []rune) int {
	srcLen, targetLen := len(src), len(target)
	column := make([]int, srcLen + 1)

	for y := 1; y <= srcLen; y++ {
		column[y] = y
	}

	for x := 1; x <= targetLen; x++ {
		column[0] = x
		lastkey := x - 1
		for y := 1; y <= srcLen; y++ {
			oldkey := column[y]
			var incr int
			if src[y-1] != target[x-1] {
				incr = 1
			}
			column[y] = minimum(column[y]+1, column[y-1]+1, lastkey+incr)
			lastkey = oldkey 
		}
	}
	return column[srcLen]
}

func minimum(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
	} else {
		if b < c {
			return b
		}
	}
	return c
}

func TestLevenshtein() {
		var str1 = []rune("Asheville")
		var str2 = []rune("Arizona")
		fmt.Println("Distance between Asheville and Arizona:",LevenshteinDistance(str1,str2))
		
		str1 = []rune("Python")
		str2 = []rune("Peithen")
		fmt.Println("Distance between Python and Peithen:",LevenshteinDistance(str1,str2))
		
		str1 = []rune("Orange")
		str2 = []rune("Apple")
		fmt.Println("Distance between Orange and Apple:",LevenshteinDistance(str1,str2))
}