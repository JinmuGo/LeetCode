
type Edge struct {
    x int
    y int
    d int
}

func nearestExit(maze [][]byte, entrance []int) int {
    n := len(maze)
    m := len(maze[0])

    ex, ey := entrance[1], entrance[0]

    maze[ey][ex] = '+'

    queue := make([]Edge, 0, n * m)
    queue = append(queue, Edge{x: ex, y: ey, d: 0})

    dx := []int{0, 0, 1, -1}
    dy := []int{1, -1, 0, 0}
    
    for len(queue) > 0 {
        elem := queue[0]
        queue = queue[1:]

        for i := 0; i < 4; i++ {
            x := dx[i] + elem.x
            y := dy[i] + elem.y

            if  x >= 0 && x < m && y >= 0 && y < n && maze[y][x] == '.' {
                
                if x == 0 || x == m - 1 || y == 0 || y == n - 1 {
                    return elem.d + 1
                }
                
                maze[y][x] = '+'
                queue = append(queue, Edge{x: x, y: y, d: elem.d + 1})
            }
        }
    }

    return -1
}