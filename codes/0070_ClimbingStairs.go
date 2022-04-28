package algorithmn

func climbStairs(n int) int {
    if n <= 2 {
        return n
    }
    
    var arr []int
    for i := 0; i < n; i++ {
        if i < 2 {
            arr = append(arr, i + 1)
        } else {
            arr = append(arr, arr[i-1] + arr[i-2])
        }
    }
    return arr[n-1]
}