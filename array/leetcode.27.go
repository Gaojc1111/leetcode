package array

func removeElement(nums []int, val int) int {
	cnt := 0
	for _, v := range nums {
		if v != val {
			nums[cnt] = v
			cnt++
		}
	}
	return cnt
}
