package leetcode

// lc136
func singleNumber(nums []int) int {
	answer := 0
	for _, v := range nums {
		answer ^= v
	}
	return answer
}
