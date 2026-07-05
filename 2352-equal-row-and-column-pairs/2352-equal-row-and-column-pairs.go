func equalPairs(grid [][]int) int {
    n := len(grid)
    rowCount := make(map[string]int)

    for r := 0; r < n; r++ {
        var sb strings.Builder
        for c := 0; c < n; c++ {
            sb.WriteString(strconv.Itoa(grid[r][c]))
            sb.WriteString(",")
        }
        rowCount[sb.String()]++
    }

    count := 0

    for c := 0; c < n; c++ {
        var sb strings.Builder
        for r := 0; r < n; r++ {
            sb.WriteString(strconv.Itoa(grid[r][c]))
            sb.WriteString(",")
        }

        count += rowCount[sb.String()]
    }
    
    return count
}