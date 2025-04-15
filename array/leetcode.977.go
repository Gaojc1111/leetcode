package array

// 双指针
func sortedSquares(nums []int) []int {
	i, j := -1, 0
	for _, v := range nums {
		if v < 0 {
			i++
		} else {
			break
		}
		j++
	}
	length := len(nums)
	res := []int{}
	for i >= 0 && j < length {
		i2 := nums[i] * nums[i]
		j2 := nums[j] * nums[j]
		if i2 < j2 {
			i--
			res = append(res, i2)
		} else {
			j++
			res = append(res, j2)
		}
	}

	for i >= 0 {
		res = append(res, nums[i]*nums[i])
		i--
	}
	for j < length {
		res = append(res, nums[j]*nums[j])
		j++
	}
	return res
}
