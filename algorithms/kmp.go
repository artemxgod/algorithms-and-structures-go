package algorithms

import (
	"fmt"
)

const patternSize = 100

// The Knuth–Morris–Pratt string searching algorithm (or KMP algorithm)
// vsearches for occurrences of a "substring" within a main "string"
// by employing the observation that when a mismatch occurs,
// the word itself embodies sufficient information to determine where
// the next match could begin, thus bypassing re-examination of previously matched characters.
// Here is source code of the Go Program to implement Knuth–Morris–Pratt (KMP) Algorithm
func KMPSearch(str, substr string) int {
	resSlice := KMP(str, substr)
	if len(resSlice) > 0 {
		return resSlice[0]
	}
	return -1
}

// the last match
func KMPSearchLast(str, substr string) int {
	resSlice := KMP(str, substr)
	if len(resSlice) > 0 {
		return resSlice[len(resSlice)-1]
	}
	return -1
}

func KMP(str, substr string) []int {
	var res []int
	next := preKMP(substr)
	strLen, subLen := len(str), len(substr)

	// no possible matches
	if strLen == 0 || subLen == 0 || strLen < subLen {
		return res
	}

	j := 0
	for i := 0; i < strLen; {
		for j > -1 && str[i] != substr[j] {
			j = next[j]
		}
		i++
		j++

		// if j reached the value of subLen => match was found
		// Match starts from i-j, where i is the index where match ends
		if j >= subLen {
			res = append(res, i-j)
			j = next[j]
		}
	}

	return res
}

func preKMP(str string) [patternSize]int {
	var KMPNext [patternSize]int
	KMPNext[0] = -1

	j := -1
	for i := 0; i < len(str)-1; {
		for j > -1 && str[i] != str[j] {
			j = KMPNext[j]
		}
		i++
		j++

		if str[i] == str[j] {
			KMPNext[i] = KMPNext[j]
		} else {
			KMPNext[i] = j
		}
	}
	return KMPNext
}

func TestKMPSearch() {
	fmt.Println("Search First Position String:")
	fmt.Printf("%2d\n", KMPSearch("cocacola", "co"))
	fmt.Printf("%2d\n", KMPSearch("Australia", "lia"))
	fmt.Printf("%2d\n", KMPSearch("cocacola", "cx"))
	fmt.Printf("%2d\n", KMPSearch("AABAACAADAABAABA", "AABA"))

	fmt.Println("Search Last Position String:")
	fmt.Printf("%2d\n", KMPSearchLast("cocacolaco", "co"))
	fmt.Printf("%2d\n", KMPSearchLast("Australia", "lia"))
	fmt.Printf("%2d\n", KMPSearchLast("cocacola", "cx"))
	fmt.Printf("%2d\n", KMPSearchLast("AABAACAADAABAABAAABAACAADAABAABA", "AABA"))
}
