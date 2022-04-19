package algorithmn

func moveZeroes(nums []int)  {
    right := len(nums) - 1
    left := right - 1
    for left >= 0 {
        if nums[left] == 0 {
            for i := left; i < right; i++ {
                nums[i], nums[i+1] = nums[i+1], nums[i]
            }
            right--
        }
        left--
    }
}