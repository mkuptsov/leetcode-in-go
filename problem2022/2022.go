package problem2022

func construct2DArray(original []int, m int, n int) [][]int {
	if len(original) == m*n {
		new := make([][]int, m)
		for i := 0; i < m; i++ {
			new[i] = original[i*n : (i+1)*n]
		}
		return new
	}
	return [][]int{}
}
