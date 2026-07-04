type Edge struct {
    Dest int
    Distance int
}

func minScore(n int, roads [][]int) int {
// roads 의 distance중에서 가장 최솟값부터 시작해서 
// 그 최솟값의 도로를 1 -> n 까지 가는 중에 갈 수 있으면 최솟값 리턴. 
    pathMap := make([][]Edge, n + 1)

    for _, road := range roads {
        u, v, d := road[0], road[1], road[2]
        pathMap[u] = append(pathMap[u], Edge{Dest: v, Distance: d})
        pathMap[v] = append(pathMap[v], Edge{Dest: u, Distance: d})
    }

    visited := make([]bool, n + 1)
    queue := []int{1}
    visited[1] = true

    minDist := math.MaxInt32

    for len(queue) > 0 {
        cur := queue[0]
        queue = queue[1:]

        for _, edge := range pathMap[cur] {
            if edge.Distance < minDist {
                minDist = edge.Distance
            }

            if !visited[edge.Dest] {
                visited[edge.Dest] = true
                queue = append(queue, edge.Dest)
            }
        }
    } 

    return minDist
}