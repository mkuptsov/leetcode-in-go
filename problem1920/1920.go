package problem1920

func buildArray(nums []int) []int {
	ans := make([]int, len(nums))
	for _, i := range nums {
		ans[i] = nums[nums[i]]
	}
	return ans
}
