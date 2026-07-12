func countCompleteComponents(n int, edges [][]int) int {
	adj := make([][]int, n)
	deg := make([]int, n)
	visited := make([]bool, n)

	for _, edge := range edges {
		a, b := edge[0], edge[1]
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
		deg[a]++
		deg[b]++
	}

	ans := 0

	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}

		queue := make([]int, 0)
		queue = append(queue, i)
		visited[i] = true

        nodeCount := 0
        edgeCount := 0

		for len(queue) > 0 {
			elem := queue[0]
			queue = queue[1:]

            nodeCount++
            edgeCount += deg[elem]

            for _, next := range adj[elem] {
                if !visited[next] {
                    visited[next] = true
                    queue = append(queue, next)
                }
            }
		}

        if edgeCount == nodeCount * (nodeCount - 1) {
            ans++
        }
	}

    return ans
}