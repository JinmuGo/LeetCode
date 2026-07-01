type cell struct {
	x, y, val int
}

type maxHeap []cell

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i].val > h[j].val } // 최대힙: val 큰 게 먼저
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x any)        { *h = append(*h, x.(cell)) }
func (h *maxHeap) Pop() any {
	old := *h
	m := len(old)
	item := old[m-1]
	*h = old[:m-1]
	return item
}

func maximumSafenessFactor(grid [][]int) int {
	n := len(grid)

	dx := []int{1, -1, 0, 0}
	dy := []int{0, 0, 1, -1}

	// 1단계: 모든 도둑을 한 번에 큐에 넣는 멀티소스 BFS.
	// 각 칸의 값 = 가장 가까운 도둑까지의 거리(= 맨해튼 거리).
	safenessMap := make([][]int, n)
	for i := range safenessMap {
		safenessMap[i] = make([]int, n)
		for j := range safenessMap[i] {
			safenessMap[i][j] = -1 // 미방문
		}
	}

	queue := [][]int{}
	for i, row := range grid {
		for j, val := range row {
			if val == 1 {
				safenessMap[i][j] = 0
				queue = append(queue, []int{j, i}) // {x=col, y=row}
			}
		}
	}

	for len(queue) > 0 {
		rx, ry := queue[0][0], queue[0][1]
		queue = queue[1:]

		for i := 0; i < 4; i++ {
			x, y := dx[i]+rx, dy[i]+ry
			if x >= 0 && x < n && y >= 0 && y < n && safenessMap[y][x] == -1 {
				safenessMap[y][x] = safenessMap[ry][rx] + 1
				queue = append(queue, []int{x, y})
			}
		}
	}

	// 2단계: 경로 위 칸들의 최솟값(=경로 안전도)을 최대화하는 경로 찾기.
	// 최대힙 다익스트라: 항상 "지금까지 안전도가 가장 높은" 칸을 먼저 확장한다.
	best := make([][]int, n)
	for i := range best {
		best[i] = make([]int, n)
		for j := range best[i] {
			best[i][j] = -1
		}
	}

	pq := &maxHeap{}
	heap.Push(pq, cell{0, 0, safenessMap[0][0]})
	best[0][0] = safenessMap[0][0]

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(cell)

		// 최대힙이므로 도착지를 처음 꺼내는 순간이 곧 최적해다.
		if cur.x == n-1 && cur.y == n-1 {
			return cur.val
		}
		// 이미 더 좋은 값으로 방문된 칸이면 스킵 (stale entry)
		if cur.val < best[cur.y][cur.x] {
			continue
		}

		for i := 0; i < 4; i++ {
			x, y := dx[i]+cur.x, dy[i]+cur.y
			if x >= 0 && x < n && y >= 0 && y < n {
				nv := min(cur.val, safenessMap[y][x]) // 한 경로 안에서는 min 누적
				if nv > best[y][x] {                  // 개선될 때만 push
					best[y][x] = nv
					heap.Push(pq, cell{x, y, nv})
				}
			}
		}
	}

	return 0
}