package patternSearchingAlgo

const base = 16777619

func RabinKarpAlgo(txt string, patterns []string) []string {
	in := indices(txt, patterns)
	matches := make([]string, len(in))
	i := 0
	for j, p := range patterns {
		if _, ok := in[j]; ok {
			matches[i] = p
			i++
		}
	}

	return matches
}

func indices(txt string, patterns []string) map[int]int {
	n, m := len(txt), minLen(patterns)
	matches := make(map[int]int)

	if n < m || len(patterns) == 0 {
		return matches
	}

	var mult uint32 = 1 // mult = base^(m-1)
	for i := 0; i < m-1; i++ {
		mult = (mult * base)
	}

	hp := hashPatterns(patterns, m)
	h := hash(txt[:m])
	for i := 0; i < n-m+1 && len(hp) > 0; i++ {
		if i > 0 {
			h = h - mult*uint32(txt[i-1])
			h = h*base + uint32(txt[i+m-1])
		}

		if mps, ok := hp[h]; ok {
			for _, pi := range mps {
				pat := patterns[pi]
				e := i + len(pat)
				if _, ok := matches[pi]; !ok && e <= n && pat == txt[i:e] {
					matches[pi] = i
				}
			}
		}
	}
	return matches
}

func hash(s string) uint32 {
	var h uint32
	for i := 0; i < len(s); i++ {
		h = (h*base + uint32(s[i]))
	}
	return h
}

func hashPatterns(patterns []string, l int) map[uint32][]int {
	m := make(map[uint32][]int)
	for i, t := range patterns {
		h := hash(t[:l])
		if _, ok := m[h]; ok {
			m[h] = append(m[h], i)
		} else {
			m[h] = []int{i}
		}
	}

	return m
}

func minLen(patterns []string) int {
	if len(patterns) == 0 {
		return 0
	}

	m := len(patterns[0])
	for i := range patterns {
		if m > len(patterns[i]) {
			m = len(patterns[i])
		}
	}

	return m
}
func LevenshteinDistance(s1, s2 string) int {
	s1len := len(s1)
	s2len := len(s2)
	column := make([]int, len(s1)+1)

	for y := 1; y <= s1len; y++ {
		column[y] = y
	}
	for x := 1; x <= s2len; x++ {
		column[0] = x
		lastkey := x - 1
		for y := 1; y <= s1len; y++ {
			oldkey := column[y]
			var incr int
			if s1[y-1] != s2[x-1] {
				incr = 1
			}

			column[y] = minimum(column[y]+1, column[y-1]+1, lastkey+incr)
			lastkey = oldkey
		}
	}
	return column[s1len]
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

func LCS(a, b string, m, n int) int {
	if m == 0 || n == 0 {
		return 0
	}

	if a[m-1] == b[n-1] {
		return 1 + LCS(a, b, m-1, n-1)
	}
	if LCS(a, b, m, n-1) > LCS(a, b, m-1, n) {
		return LCS(a, b, m, n-1)
	} else {
		return LCS(a, b, m-1, n)
	}
}

func SearchNext(haystack string, needle string) int {
	retSlice := KMP(haystack, needle)
	if len(retSlice) > 0 {
		return retSlice[len(retSlice)-1]
	}
	return -1
}

func KMP_SearchString(haystack string, needle string) int {
	retSlice := KMP(haystack, needle)
	if len(retSlice) > 0 {
		return retSlice[0]
	}

	return -1
}

const PatternSize = 256

func KMP(haystack string, needle string) []int {
	next := preKMP(needle)
	i := 0
	j := 0
	m := len(needle)
	n := len(haystack)

	x := []byte(needle)
	y := []byte(haystack)
	var ret []int

	//got zero target or want, just return empty result
	if m == 0 || n == 0 {
		return ret
	}

	//want string bigger than target string
	if n < m {
		return ret
	}

	for j < n {
		for i > -1 && x[i] != y[j] {
			i = next[i]
		}
		i++
		j++

		//fmt.Println(i, j)
		if i >= m {
			ret = append(ret, j-i)
			//fmt.Println("find:", j, i)
			i = next[i]
		}
	}

	return ret
}

func preMP(x string) [PatternSize]int {
	var i, j int
	length := len(x) - 1
	var mpNext [PatternSize]int
	i = 0
	j = -1
	mpNext[0] = -1

	for i < length {
		for j > -1 && x[i] != x[j] {
			j = mpNext[j]
		}
		i++
		j++
		mpNext[i] = j
	}
	return mpNext
}

func preKMP(x string) [PatternSize]int {
	var i, j int
	length := len(x) - 1
	var kmpNext [PatternSize]int
	i = 0
	j = -1
	kmpNext[0] = -1

	for i < length {
		for j > -1 && x[i] != x[j] {
			j = kmpNext[j]
		}

		i++
		j++

		if x[i] == x[j] {
			kmpNext[i] = kmpNext[j]
		} else {
			kmpNext[i] = j
		}
	}
	return kmpNext
}
