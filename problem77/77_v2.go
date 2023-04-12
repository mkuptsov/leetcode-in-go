package problem77

func combineV2(n int, k int) [][]int {
	var result [][]int
	buf := make([]int, 0, k)

	var backtrack func(int, int, []int)

	backtrack = func(start int, left int, cur []int) {
		if left == 0 {
			curCopy := make([]int, len(cur))
			copy(curCopy, cur)
			result = append(result, curCopy)
			return
		}

		for i := start; i <= n; i++ {
			backtrack(i+1, left-1, append(cur, i))
		}
	}
	backtrack(1, k, buf)

	return result
}
