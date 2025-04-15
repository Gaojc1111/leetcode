package array

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	i, j := 0, -1
	cnt := 1
	for cnt <= n*n {
		// 右
		for j+1 < n {
			if res[i][j+1] == 0 {
				res[i][j+1] = cnt
				cnt++
				j++
			} else {
				break
			}
		}

		// 下
		for i+1 < n {
			if res[i+1][j] == 0 {
				res[i+1][j] = cnt
				cnt++
				i++
			} else {
				break
			}
		}

		// 左
		for j-1 >= 0 {
			if res[i][j-1] == 0 {
				res[i][j-1] = cnt
				cnt++
				j--
			} else {
				break
			}
		}

		// 上
		for i-1 >= 0 {
			if res[i-1][j] == 0 {
				res[i-1][j] = cnt
				cnt++
				i--
			} else {
				break
			}
		}
	}
	return res
}
