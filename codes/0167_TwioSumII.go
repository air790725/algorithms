package algorithmn

func twoSum(numbers []int, target int) []int {
    for left, right := 0, len(numbers) - 1; left < right; {
        total := numbers[left] + numbers[right]
        if total == target {
            return []int{left+1, right+1}
        }
        if total < target {
            left++
        } else {
            right--
        }
    }
    return make([]int, 2)
}