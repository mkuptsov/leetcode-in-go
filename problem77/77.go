package problem77

func combine(n int, k int) [][]int {
	var result [][]int
	buf := make([]int, 0, k)

	var backtrack func(int, int)

	backtrack = func(start int, left int) {
		if left == 0 {
			bufCopy := make([]int, len(buf))
			copy(bufCopy, buf)
			result = append(result, bufCopy)
			return
		}

		for i := start; i <= n; i++ {
			buf = append(buf, i)
			backtrack(i+1, left-1)
			buf = buf[:len(buf)-1]
		}
	}
	backtrack(1, k)

	return result
}
