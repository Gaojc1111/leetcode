package hash

// 思路：1. 字符串长度相等 2. 字母类型总数相等 3. 每个字母出现次数相等
func isAnagram(s string, t string) bool {
	lenS := len(s)
	lenT := len(t)
	if lenS != lenT {
		return false
	}

	m1 := map[uint8]int{} // unicode: uint8 -> rune
	m2 := map[uint8]int{}

	for i := 0; i < lenS; i++ {
		m1[s[i]]++
		m2[t[i]]++
	}

	if len(m1) != len(m2) {
		return false
	}

	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}
	return true
}
