package problem1

func TwoSum(nums []int, target int) []int {
    var index1, index2 int
    for i := 0; i < len(nums) - 1; i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[i] + nums[j] == target {
                index1, index2 = i, j
            }
        }
    }
    return []int{index1, index2}
}
