package algorithmn

func rotate(nums []int, k int)  {
    length := len(nums)
    for k > 0 {
        last := nums[length-1]
        for i := length-1; i > 0 ; i-- {
            nums[i] = nums[i-1]
        }
        nums[0] = last
        k--
    }
}