package algorithmn

func sortedSquares(nums []int) []int {
    left, right, result := 0, len(nums)-1, make([]int, len(nums))
    for index := right; index >= 0; index-- {
        a := nums[left] * nums[left]
        b := nums[right] * nums[right]
        if a <= b {
            right--
            result[index] = b
        } else {
            left++
            result[index] = a
        }
    }
    return result
}