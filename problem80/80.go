package problem80

func RemoveDuplicates(nums []int) int {
    counter := 1
    j := 0
    for i := 1; i < len(nums); i++ {
        if nums[i] != nums[j] {
            j++
            nums[j] = nums[i]
            counter = 1
        } else if counter < 2 {
            j++
            nums[j] = nums[i]
            counter++
        }
    }
    return j + 1
}
