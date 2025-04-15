package array

func minSubArrayLen(target int, nums []int) int {
	i, j := 0, 0
	length := len(nums)
	sum := 0
	ans := 100010
	for j < length {
		sum += nums[j]

		for sum >= target {
			ans = min(ans, j-i+1)
			sum -= nums[i]
			i++
		}
		j++
	}
	if ans == 100010 {
		ans = 0
	}
	return ans
}
