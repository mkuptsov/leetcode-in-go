package problem78

func subsets(nums []int) [][]int {
	result := [][]int{{}}
	buf := make([]int, 0, len(nums))

	var backtrack func(int, int, []int)
	backtrack = func(start, left int, cur []int) {
		if left == 0 {
			return
		}

		for i := start; i < len(nums); i++ {
			buf = append(buf, cur[i])
			copied := make([]int, len(buf))
			copy(copied, buf)
			result = append(result, copied)
			backtrack(i+1, left-1, cur)
			buf = buf[:len(buf)-1]
		}
	}
	backtrack(0, len(nums), nums)
	return result
}
