func findCircleNum(isConnected [][]int) int {
    n := len(isConnected)

    // isConnected[i][j] == isConnected[j][i]. important!

    visited := make([]bool, n)
    res := 0

    for city := 0; city < n; city++ {
        if visited[city] == true {
            continue
        }
        
        visited[city] = true
        res++
        queue := make([]int, 0)
        queue = append(queue, city)

        for len(queue) > 0 {
            elem := queue[0]
            queue = queue[1:]

            for i := 0; i < n; i++ {
                if isConnected[elem][i] == 1 && !visited[i] {
                    visited[i] = true
                    queue = append(queue, i)
                }
            }
        }
    }

    return res
}