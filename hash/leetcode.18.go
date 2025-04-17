package hash

import "sort"

func fourSum(nums []int, target int) [][]int {
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
			for k := j + 1; k < length; k++ {
				if k != j+1 && nums[k] == nums[k-1] {
					continue
				}
				t := target - nums[i] - nums[j] - nums[k]
				if m[t] > k {
					ans = append(ans, []int{nums[i], nums[j], nums[k], t})
				}
			}
		}
	}
	return ans
}
