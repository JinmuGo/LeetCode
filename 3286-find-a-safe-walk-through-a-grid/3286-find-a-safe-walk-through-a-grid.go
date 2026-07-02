func findSafeWalk(grid [][]int, health int) bool {
    m := len(grid)
    n := len(grid[0])
    // health >= 1 && grid[n -1][n - 1] 에 도착
    // 다 돌면서 맵을 만들어야 겠는데. 현재 위치에 도달하면 체력이 어느정도 남았는지. 만약 n- 1, m-1 ?이 1 >= 이상이면 true 아니면  false나오도록. 

    healthMap := make([][]int, m)
    queue := make([][]int, 0)

    for i := range healthMap {
        healthMap[i] = make([]int, n)
    }

    dx := []int{0, 0, 1, -1}
    dy := []int{1, -1, 0, 0}

    queue = append(queue, []int{0, 0, health - grid[0][0]})

    for len(queue) > 0 {
        elem := queue[0]
        queue = queue[1:]
        curX := elem[0]
        curY := elem[1]
        h := elem[2]

        if h < healthMap[curY][curX] {
            continue
        }

        for i := 0; i < 4; i++ {
            x := curX + dx[i]
            y := curY + dy[i]
            
            if x >= 0 && n > x && y >= 0 && m > y {
                nextH := h - grid[y][x]

                if nextH > healthMap[y][x] && nextH > 0 {
                    healthMap[y][x] = nextH
                    queue = append(queue, []int{x, y, nextH})
                }
            }
        }
    }

    return healthMap[m - 1][n - 1] >= 1
}