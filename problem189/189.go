package problem189

func rotate(nums []int, k int) {
	if len(nums) < 2 || k == 0 {
		return
	}
	r := len(nums) - k%len(nums)
	copy(nums, append(nums[r:], nums[:r]...))
}
