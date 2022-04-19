package algorithmn

func rotate(nums []int, k int) {
    k %= len(nums)
    reverse(nums)
    reverse(nums[:k])
    reverse(nums[k:])
}

func reverse(arr []int) {
    for i, n := 0, len(arr); i < n/2; i++ {
        arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
    }
}