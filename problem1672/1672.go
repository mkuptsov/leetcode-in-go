package problem1672

func maximumWealth(accounts [][]int) int {
	var max int
	for _, account := range accounts {
		var sum int
		for _, wealth := range account {
			sum += wealth
		}
		if sum > max {
			max = sum
		}
	}
	return max
}
