package array

func search(nums []int, target int) int {
	return erfen(nums, target)
}

func erfen(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r + 1) >> 1
		if nums[mid] <= target {
			l = mid
		} else {
			r = mid - 1
		}
	}
	if nums[l] != target {
		return -1
	}
	return l
}
