type edge struct {
    to int
    dir bool
}

func minReorder(n int, connections [][]int) int {
    graph := make([][]edge, n)

    for _, con := range connections {
        from, to := con[0], con[1]

        graph[from] = append(graph[from], edge{to: to, dir: true})
        graph[to] = append(graph[to], edge{to: from, dir: false})
    }

    cnt := 0
    var dfs func(from, parent int) 

    dfs = func(from, parent int)  {
        for _, edge := range graph[from] {
            if edge.to == parent {
                continue
            }

            if edge.dir {
                cnt++
            }
            dfs(edge.to, from)
        }
    }

    dfs(0, -1)

    return cnt
}