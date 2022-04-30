package algorithmn

func minimumTotal(triangle [][]int) int {
    for row := len(triangle) - 2; row >= 0; row-- {
        for col := 0; col < len(triangle[row]); col++ {
            if triangle[row+1][col] < triangle[row+1][col+1] {
                triangle[row][col] += triangle[row+1][col]
            } else {
                triangle[row][col] += triangle[row+1][col+1]
            }   
        }
    }
    return triangle[0][0]
}