package problem1342

func numberOfSteps(num int) int {
	i := 0
	for num > 0 {
		if num&1 == 0 {
			num >>= 1
		} else {
			num--
		}
		i++
	}
	return i
}
