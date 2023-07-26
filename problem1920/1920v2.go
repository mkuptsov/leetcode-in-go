package problem1920

func buildArrayV2(nums []int) []int {
	for i, n := range nums {
		nums[i] += (nums[n] % 1000) * 1000
	}

	for i, n := range nums {
		nums[i] = n / 1000
	}
	return nums
}
