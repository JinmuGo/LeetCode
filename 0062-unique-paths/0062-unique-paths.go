func uniquePaths(m int, n int) int {
    grid := make([][]int, m)

    for x := range grid {
        grid[x] = make([]int, n)
    }

    grid[0][0] = 1

    for i := range m {
        for j := range n {
            if i == 0 || j == 0 {
                grid[i][j] = 1
                continue 
            }
            grid[i][j] = grid[i - 1][j] + grid[i][j - 1]
        }
    }
    
    return grid[m - 1][n - 1]
}