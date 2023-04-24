package problem119

func getRow(rowIndex int) []int {
	var result []int
	for i := 0; i <= rowIndex; i++ {
		result = append(result, 1)
		result[0], result[i] = 1, 1
		if i > 1 {
			for j := i - 1; j > 0; j-- {
				result[j] = result[j] + result[j-1]
			}
		}
	}
	return result
}
