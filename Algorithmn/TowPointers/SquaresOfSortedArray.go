package algorithmn

import (
    "math"
)

func sortedSquares(nums []int) []int {
    left, right := 0, len(nums)-1
    index, array := right, make([]int, right+1)
    for index >= 0 {
        a := math.Pow(float64(nums[left]), 2)
        b := math.Pow(float64(nums[right]), 2)
        if a <= b {
            right--
            array[index] = int(b)
        } else {
            left++
            array[index] = int(a)
        }
        index--
    }
    return array
}