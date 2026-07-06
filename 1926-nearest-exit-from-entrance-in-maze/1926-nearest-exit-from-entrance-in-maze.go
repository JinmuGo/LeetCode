
type Edge struct {
    x int
    y int
    d int
}

func nearestExit(maze [][]byte, entrance []int) int {
    n := len(maze)
    m := len(maze[0])
    // empty cells = '.' , wall = '+'  
    // entrance => 처음 진입점

    visited := make([][]bool, n)

    for i := range visited {
        visited[i] = make([]bool, m)
    }

    ex := entrance[1]
    ey := entrance[0]

    visited[ey][ex] = true

    queue := make([]Edge, 0)

    queue = append(queue, Edge{x: ex, y: ey, d: 0})

    dx := []int{0, 0, 1, -1}
    dy := []int{1, -1, 0, 0}

    res := math.MaxInt32
    
    for len(queue) > 0 {
        elem := queue[0]
        queue = queue[1:]

        for i := 0; i < 4; i++ {
            x := dx[i] + elem.x
            y := dy[i] + elem.y

            if  x >= 0 && x < m && y >= 0 && y < n && !visited[y][x] {
                visited[y][x] = true

                if maze[y][x] == '+' {
                    continue 
                }

                queue = append(queue, Edge{x: x, y: y, d: elem.d + 1})
                
                if x == 0 || x == m - 1 || y == 0 || y == n - 1 {
                    val := elem.d + 1

                    if res > val {
                        res = val
                    }
                }
                
            }
        }
    }

    if res == math.MaxInt32 {
        return -1
    }

    return res
}