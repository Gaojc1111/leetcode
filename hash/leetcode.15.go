package hash

import "sort"

// 思路: 先排序，确定两个数，第三个数下标大于第二个数，最后避免重复，剪枝。
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	m := map[int]int{}
	for k, v := range nums {
		m[v] = k
	}

	ans := [][]int{}
	length := len(nums)
	for i := 0; i < length; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < length; j++ {
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}
			if m[-nums[i]-nums[j]] > j {
				ans = append(ans, []int{nums[i], nums[j], -nums[i] - nums[j]})
			}
		}
	}
	return ans
}
