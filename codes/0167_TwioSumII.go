package algorithmn

func twoSum(numbers []int, target int) []int {
    left, right, result := 0, len(numbers) - 1, make([]int, 2)
    for left < right {
        total := numbers[left] + numbers[right]
        if total < target {
            left++
        } else if total == target {
            result[0] = left + 1
            result[1] = right + 1
            return result
        } else {
            right--
        }
    }
    return result
}