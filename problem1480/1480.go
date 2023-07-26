package problem1480

func runningSum(nums []int) []int {
	runningSum := make([]int, len(nums))
	runningSum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		runningSum[i] = runningSum[i-1] + nums[i]
	}
	return runningSum
}
