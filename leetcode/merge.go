package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	i1 := m - 1
	i2 := n - 1
	merged := m + n - 1

	for i1 >= 0 || i2 >= 0 {
		if i1 < 0 {
			nums1[merged] = nums2[i2]
			i2--
		} else if i2 < 0 {
			nums1[merged] = nums1[i1]
			i1--
		} else if nums1[i1] < nums2[i2] {
			nums1[merged] = nums2[i2]
			i2--
		} else {
			nums1[merged] = nums1[i1]
			i1--
		}
		merged--
	}
}
