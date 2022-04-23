package algorithmn

func moveZeroes(nums []int)  {
    for left, right := len(nums) - 2, len(nums) - 1; left >= 0; left-- {
        if nums[left] == 0 {
            for i := left; i < right; i++ {
                nums[i], nums[i+1] = nums[i+1], nums[i]
            }
            right--
        }
    }
}