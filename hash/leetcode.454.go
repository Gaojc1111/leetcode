package hash

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	m1 := map[int]int{}
	m2 := map[int]int{}

	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			m1[v1+v2]++
		}
	}
	for _, v1 := range nums3 {
		for _, v2 := range nums4 {
			m2[v1+v2]++
		}
	}

	ans := 0
	for k, _ := range m1 {
		ans += m1[k] * m2[-k]
	}
	return ans
}
