package algorithms

const base = 16777619

// The Rabin-Karp algorithm is a string searching algorithm that uses hashing to find patterns in strings.
// Here is source code of the Go Program to Implement Rabin-Karp string search algorithm.
func RKSearch(text string, patterns []string) []string {
	matchesIDX := indices(text, patterns)
	matches := make([]string, len(matchesIDX))
	i := 0
	for jdx, ptrn := range patterns {
		if _, ok := matchesIDX[jdx]; ok {
			// fmt.Println(string(text[index:len(ptrn)+index]))
			matches[i] = ptrn
			i++
		}
	}
	return matches
}

func indices(txt string, patterns []string) map[int]int {
	n, m := len(txt), minLen(patterns)
	matches := make(map[int]int)

	// check if matches can be found
	if n < m || len(patterns) == 0 {
		return matches
	}

	// mult = base^(m-1)
	var mult uint32 = 1
	for i := 0; i < m-1; i++ {
		mult *= base
	}

	hashedPatterns := hashPatterns(patterns, m)
	hashedText := hash(txt[:m])

	for i := 0; i < n-m+1 && len(hashedPatterns) > 0; i++ {
		// first m symbols are hashed so no need to move further in the first iteration 
		if i > 0 {
			hashedText = hashedText - mult*uint32(txt[i-1]) // remove the first elem of substring
			hashedText = hashedText*base + uint32(txt[i+m-1]) // add next elem of substring
		}

		if matchedPatterns, ok := hashedPatterns[hashedText]; ok {
			for _, pi := range matchedPatterns {
				pat := patterns[pi] // found pattern
				matchLen := i + len(pat) // a found pattern len
				// here we compare pattern and substring
				if _, ok := matches[pi]; !ok && matchLen <= n && pat == txt[i:matchLen] {
					matches[pi] = i
				}
			}
		}
	}
	return matches
}

// hash a text
func hash(s string) uint32 {
	var h uint32
	for i := 0; i < len(s); i++ {
		h = h*base + uint32(s[i])
	}
	return h
}

// hash a pattern into map, where hash is key and value is an index
func hashPatterns(patterns []string, l int) map[uint32][]int {
	hashPatternsMap := make(map[uint32][]int)
	for idx, pattern := range patterns {
		h := hash(pattern[:l])
		hashPatternsMap[h] = append(hashPatternsMap[h], idx)
	}

	return hashPatternsMap
}

// finds min len of patterns
func minLen(patterns []string) int {
	if len(patterns) == 0 {
		return 0
	}

	min := len(patterns[0])
	for idx := range patterns {
		if len(patterns[idx]) < min {
			min = len(patterns[idx])
		}
	}
	return min
}
