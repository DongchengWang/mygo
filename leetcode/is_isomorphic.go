package leetcode

// lc205
func isIsomorphic(s string, t string) bool {
	var prevIndexOfS [128]int
	var prevIndexOfT [128]int
	for i := 0; i < len(s); i++ {
		sc := s[i]
		tc := t[i]
		if prevIndexOfS[sc] != prevIndexOfT[tc] {
			return false
		}
		// Save initial index 0
		prevIndexOfS[sc] = i + 1
		prevIndexOfT[tc] = i + 1
	}
	return true
}
