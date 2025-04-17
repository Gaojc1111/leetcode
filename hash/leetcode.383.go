package hash

func canConstruct(ransomNote string, magazine string) bool {
	arr := [26]int{}
	for _, v := range magazine {
		arr[v-'a']++
	}
	for _, v := range ransomNote {
		arr[v-'a']--
		if arr[v-'a'] < 0 {
			return false
		}
	}
	return true
}
