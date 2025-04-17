package hash

func twoSum(nums []int, target int) []int {
	m := map[int]int{}

	for k, v := range nums {
		if i, ok := m[target-v]; ok {
			return []int{i, k}
		}
		m[v] = k
	}
	return nil
}
