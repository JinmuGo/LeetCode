type Orange struct {
	x int
	y int
	t int
}

func orangesRotting(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	visited := make([][]bool, m)

	for i := range visited {
		visited[i] = make([]bool, n)
	}

	dx := []int{0, 0, -1, 1}
	dy := []int{1, -1, 0, 0}

	queue := make([]Orange, 0)

	for y, col := range grid {
		for x, row := range col {
			if row == 2 {
				queue = append(queue, Orange{x: x, y: y, t: 0})
			}
		}
	}

	max := 0

	for len(queue) > 0 {
		elem := queue[0]
		queue = queue[1:]

		if max < elem.t {
			max = elem.t
		}

		for i := 0; i < 4; i++ {
			x := elem.x + dx[i]
			y := elem.y + dy[i]

			if 0 <= x && x < n && 0 <= y && y < m && !visited[y][x] && grid[y][x] == 1 {
				grid[y][x] = 2
				visited[y][x] = true
				t := elem.t + 1

				queue = append(queue, Orange{x: x, y: y, t: t})
			}
		}
	}

	for _, col := range grid {
		for _, row := range col {
			if row == 1 {
				return -1
			}
		}
	}

	return max
}