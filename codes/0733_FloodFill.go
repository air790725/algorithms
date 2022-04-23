package algorithmn

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
    var queue [][]int
    queue = append(queue, []int{sr, sc})
    marked := make(map[string]int)
    
    for len(queue) > 0 {
        row, col := queue[0][0], queue[0][1]
        index := string(row) + string(col)
        if _, ok := marked[index]; ok == false {
            if row > 0 && image[row-1][col] == image[row][col] {
                queue = append(queue, []int{row-1, col})
            }
            if row < len(image) -1 && image[row+1][col] == image[row][col] {
                queue = append(queue, []int{row+1, col})
            }
            if col > 0 && image[row][col-1] == image[row][col] {
                queue = append(queue, []int{row, col-1})
            }
            if col < len(image[row]) - 1 && image[row][col+1] == image[row][col] {
                queue = append(queue, []int{row, col+1})
            }
            marked[index] = 1
            image[row][col] = newColor
        }
        queue = queue[1:]
    }
    
    return image
}