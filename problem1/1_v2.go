func twoSum(nums []int, target int) []int {
    numsMap := map[int]int{}
    var result []int
    for i := 0; i < len(nums); i++ {
        if _, ok := numsMap[target - nums[i]]; ok {
            result = []int{numsMap[target - nums[i]], i}
            break
        } else {
            numsMap[nums[i]] = i
        }
    }
    return result
}