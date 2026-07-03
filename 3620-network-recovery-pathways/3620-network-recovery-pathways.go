type Edge struct {
    v       int
    cost    int64
}

func findMaxPathScore(edges [][]int, online []bool, k int64) int {
    n := len(online)

    maxCost := 0

    for _, edge := range edges {
        if edge[2] > maxCost {
            maxCost = edge[2]
        }
    }

    adj := make([][]Edge, n)
    for _, e := range edges {
        u, v, cost := e[0], e[1], int64(e[2])
        adj[u] = append(adj[u], Edge{v: v, cost: cost})
    }

    left, right := 0, maxCost
    ans := -1

    for left <= right {
        mid := left + (right - left) / 2

        if checkDAG(mid, n, adj, online, k) {
            ans = mid
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    return ans
}

func checkDAG(mid, n int, adj [][]Edge, online []bool, k int64) bool {
    if !online[0] || !online[n - 1] {
        return false
    }

    indegree := make([]int, n)
    for u := 0; u < n; u++ {
        if !online[u] {
            continue
        }

        for _, edge := range adj[u] {
            v := edge.v
            
            if online[v] && edge.cost >= int64(mid) {
                indegree[v]++
            }
        }
    }

    dp := make([]int64, n)
    for i := range dp {
        dp[i] = math.MaxInt64
    }
    dp[0] = 0

    queue := make([]int, 0, n)
    for i := 0; i < n; i++ {
        if online[i] && indegree[i] == 0 {
            queue = append(queue, i)
        }
    }

    for len(queue) > 0 {
        u := queue[0]
        queue = queue[1:]

        if dp[u] != math.MaxInt64 {
            for _, edge := range adj[u] {
                v := edge.v

                if !online[v] || edge.cost < int64(mid) {
                    continue
                }

                if dp[u] + edge.cost < dp[v] {
                    dp[v] = dp[u] + edge.cost
                }
            }
        }

        for _, edge := range adj[u] {
            v := edge.v

            if !online[v] || edge.cost < int64(mid) {
                continue
            }

            indegree[v]--

            if indegree[v] == 0 {
                queue = append(queue, v)
            }
        }
    }

    return dp[n - 1] <= k
}