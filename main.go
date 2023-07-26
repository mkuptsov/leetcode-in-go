package main

import "fmt"

func main() {
	nums := []int{0, 2, 1, 5, 3, 4}
	ans := buildArray(nums)
	fmt.Println(ans)
}

func buildArray(nums []int) []int {
	for i, n := range nums {
		nums[i] += (nums[n] % 1000) * 1000
	}

	for i := range nums {
		nums[i] = nums[i] / 1000
	}

	return nums
}
