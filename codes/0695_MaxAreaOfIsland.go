package algorithmn

func maxAreaOfIsland(grid [][]int) int {
    max := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if grid[i][j] > 0 {
                sum := walk(i, j, grid, 0)
                if sum > max {
                    max = sum
                }
            }
        }
    }
    return max
}

func walk(row int, col int, grid [][]int, sum int) int {
    if grid[row][col] > 0 {
        sum += grid[row][col]
        grid[row][col] = -1
        if row > 0 && grid[row-1][col] > 0 {
            sum = walk(row-1, col, grid, sum)
            grid[row-1][col] = -1
        }
        if row < len(grid) - 1 && grid[row+1][col] > 0 {
            sum = walk(row+1, col, grid, sum)
            grid[row+1][col] = -1
        }
        if col > 0 && grid[row][col-1] > 0 {
            sum = walk(row, col-1, grid, sum)
            grid[row][col-1] = -1
        }
        if col < len(grid[row]) - 1 && grid[row][col+1] > 0 {
            sum = walk(row, col+1, grid, sum)
            grid[row][col+1] = -1
        }
    }
    return sum
}