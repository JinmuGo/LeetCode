func maximumSafenessFactor(grid [][]int) int {
    n := len(grid)
    if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
        return 0
    }

    safeness := make([][]int, n)
    for i := range safeness {
        safeness[i] = make([]int, n)
        for j := range safeness[i] {
            safeness[i][j] = -1
        }
    }

    type pair struct{ r, c int }
    q := []pair{}

    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                q = append(q, pair{i, j})
                safeness[i][j] = 0
            }
        }
    }

    dirs := []pair{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

    for len(q) > 0 {
        curr := q[0]
        q = q[1:]
        for _, d := range dirs {
            nr, nc := curr.r+d.r, curr.c+d.c
            if nr >= 0 && nr < n && nc >= 0 && nc < n && safeness[nr][nc] == -1 {
                safeness[nr][nc] = safeness[curr.r][curr.c] + 1
                q = append(q, pair{nr, nc})
            }
        }
    }

    isValid := func(minSafe int) bool {
        if safeness[0][0] < minSafe || safeness[n-1][n-1] < minSafe {
            return false
        }
        
        visited := make([][]bool, n)
        for i := range visited {
            visited[i] = make([]bool, n)
        }
        
        bfsQ := []pair{{0, 0}}
        visited[0][0] = true
        
        for len(bfsQ) > 0 {
            curr := bfsQ[0]
            bfsQ = bfsQ[1:]
            
            if curr.r == n-1 && curr.c == n-1 {
                return true
            }
            
            for _, d := range dirs {
                nr, nc := curr.r+d.r, curr.c+d.c
                if nr >= 0 && nr < n && nc >= 0 && nc < n && !visited[nr][nc] && safeness[nr][nc] >= minSafe {
                    visited[nr][nc] = true
                    bfsQ = append(bfsQ, pair{nr, nc})
                }
            }
        }
        return false
    }

    low, high, ans := 0, n*2, 0
    for low <= high {
        mid := low + (high-low)/2
        if isValid(mid) {
            ans = mid
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    
    return ans
}