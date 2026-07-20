func shiftGrid(grid [][]int, k int) [][]int {
    n := len(grid[0])
    m := len(grid)
    total := m * n

    k = k % total
    if k == 0 {
        return grid
    }

    result := make([][]int, m)
    for i := range result {
        result[i] = make([]int, n)
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            newIdx := (i*n + j + k) % total
            
            newR := newIdx / n
            newC := newIdx % n
            
            result[newR][newC] = grid[i][j]
        }
    }

    return result
}