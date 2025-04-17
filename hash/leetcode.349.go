package hash

func intersection(nums1 []int, nums2 []int) []int {
	m1 := map[int]bool{}

	for _, v := range nums1 {
		m1[v] = true
	}

	ans := []int{}
	for _, v := range nums2 {
		if m1[v] {
			ans = append(ans, v)
			m1[v] = false
		}
	}

	return ans
}
