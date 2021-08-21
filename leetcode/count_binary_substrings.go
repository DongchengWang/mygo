package leetcode

// lc696
func countBinarySubstrings(s string) int {
	curr := 1
	prev := 0
	count := 0
	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i] {
			curr++
		} else {
			prev = curr
			curr = 1
		}

		if prev >= curr {
			count++
		}
	}
	return count
}
