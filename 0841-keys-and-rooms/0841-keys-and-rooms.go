func canVisitAllRooms(rooms [][]int) bool {
    n := len(rooms)
    visited := make(map[int]bool, n)

    var dfs func(keys []int)

    dfs = func(keys []int) {
        for _, key := range keys {
            if !visited[key] {
                visited[key] = true
                dfs(rooms[key])
            }
        }
    }

    visited[0] = true
    dfs(rooms[0])

    return len(visited) == n
}
