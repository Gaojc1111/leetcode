package review

// 思路: 因为数字的变化被锁死在范围内，所以要么能变成1循环，要么不能然后循环
func isHappy(n int) bool {
	m := map[int]bool{}
	for !m[n] {
		if n == 1 {
			return true
		}
		m[n] = true
		sum := 0
		for n != 0 {
			sum += (n % 10) * (n % 10)
			n /= 10
		}
		n = sum
	}
	return false
}
