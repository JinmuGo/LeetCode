type Orange struct {
	x int
	y int
}

func orangesRotting(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	dx := []int{0, 0, -1, 1}
	dy := []int{1, -1, 0, 0}

	queue := make([]Orange, 0)
	freshOrange := 0

	for y, col := range grid {
		for x, row := range col {
			if row == 2 {
				queue = append(queue, Orange{x: x, y: y})
			} else if row == 1 {
				freshOrange++
			}
		}
	}

	if freshOrange == 0 {
		return 0
	}

	minutes := 0

	for len(queue) > 0 && freshOrange > 0 {
		qLen := len(queue)

		for i := 0; i < qLen; i++ {
			elem := queue[0]
			queue = queue[1:]

			for j := 0; j < 4; j++ {
				x := elem.x + dx[j]
				y := elem.y + dy[j]

				if 0 <= x && x < n && 0 <= y && y < m && grid[y][x] == 1 {
					grid[y][x] = 2
                    freshOrange--
					queue = append(queue, Orange{x, y})
				}
			}
		}
        minutes++
	}

    if freshOrange > 0 {
        return -1
    }

	return minutes
}