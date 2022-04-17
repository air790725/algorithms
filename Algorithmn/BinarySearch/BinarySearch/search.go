import (
    "math"
)

func search(nums []int, target int) int {
    left, right := 0, len(nums) - 1 

    for left <= right {
        mid := int(math.Floor(float64((right - left)/2))) + left

        if nums[mid] < target {
            left = mid + 1
        } else if nums[mid] == target {
            return mid
        } else {
            right = mid - 1
        }
    }

    return -1
}