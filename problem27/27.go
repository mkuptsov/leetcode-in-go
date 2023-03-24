package problem27

func RemoveElement(nums []int, val int) int {
    index := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != val {
            nums[index] = nums[i]
            index++
        }
    }
    return index
}
