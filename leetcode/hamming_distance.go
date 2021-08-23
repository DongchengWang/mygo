package leetcode

// lc461
func hammingDistance(x int, y int) int {
	z := x ^ y
	count := 0
	for z != 0 {
		z &= z - 1
		count++
	}
	return count
}
