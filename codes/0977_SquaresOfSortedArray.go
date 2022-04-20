package algorithmn

func sortedSquares(nums []int) []int {
    left, right := 0, len(nums)-1
    index, array := right, make([]int, right+1)
    for index >= 0 {
        a := nums[left] * nums[left]
        b := nums[right] * nums[right]
        if a <= b {
            right--
            array[index] = b
        } else {
            left++
            array[index] = a
        }
        index--
    }
    return array
}