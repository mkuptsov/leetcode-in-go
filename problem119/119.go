package problem119

func getRow(rowIndex int) []int {
	result := make([][]int, 2)

	for i := 0; i <= rowIndex; i++ {
		result[1] = make([]int, i+1)
		result[1][0], result[1][i] = 1, 1
		for j := 1; j < i; j++ {
			result[1][j] = result[0][j-1] + result[0][j]
		}
		result[0] = make([]int, len(result[1]))
		copy(result[0], result[1])
	}
	return result[1]
}
