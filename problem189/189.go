package problem189

func rotate(nums []int, k int) {
	k = k % len(nums)
	if len(nums) < 2 || k == 0 {
		return
	}
	r := len(nums) - k
	copy(nums, append(nums[r:], nums[:r]...))
}
