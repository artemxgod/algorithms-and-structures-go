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
	matchIdx := 0
	for strIdx := 0; strIdx < strLen; {
		for matchIdx > -1 && str[strIdx] != substr[matchIdx] {
			matchIdx = next[matchIdx]
		}
		strIdx++
		matchIdx++

		// if j reached the value of subLen => match was found
		// Match starts from i-j, where i is the index where match ends
		if matchIdx >= subLen {
			res = append(res, strIdx-matchIdx)
			matchIdx = next[matchIdx]
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

// Some example of how it works
// txt = “AAAAABAAABA”
// pat = “AAAA”

// We compare first window of txt with pat

// txt = “AAAAABAAABA”
// pat = “AAAA”  [Initial position]
// We find a match. This is same as Naive String Matching.

// In the next step, we compare next window of txt with pat.

// txt = “AAAAABAAABA”
// pat =  “AAAA” [Pattern shifted one position]

// This is where KMP does optimization over Naive. In this second window, we only compare fourth A of pattern
// with fourth character of current window of text to decide whether current window matches or not. Since we know
// first three characters will anyway match, we skipped matching first three characters.

// Need of Preprocessing?

// An important question arises from the above explanation, how to know how many characters to be skipped. To know this,
// we pre-process pattern and prepare an integer array lps[] that tells us the count of characters to be skipped
