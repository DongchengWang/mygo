package leetcode

// lc647
func countSubstrings(s string) int {
	count := 0
	extendSubstrings := func(s string, start int, end int) {
		for start >= 0 && end < len(s) && s[start] == s[end] {
			start--
			end++
			count++
		}
	}
	for i := 0; i < len(s); i++ {
		extendSubstrings(s, i, i)   // odd
		extendSubstrings(s, i, i+1) // even
	}
	return count
}
