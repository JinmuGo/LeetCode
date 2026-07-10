type Node struct {
	idx int
	num int
}

func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []int {
	nodeNums := make([]Node, n)

	for i, num := range nums {
		nodeNums[i] = Node{idx: i, num: num}
	}

	slices.SortFunc(nodeNums, func(a, b Node) int {
		return cmp.Compare(a.num, b.num)
	})

	pos := make([]int, n)
	for i, node := range nodeNums {
		pos[node.idx] = i
	}

	farthest := make([]int, n)
	r := 0
	for i := 0; i < n; i++ {
		if r < i {
			r = i
		}
		for r+1 < n && nodeNums[r+1].num-nodeNums[i].num <= maxDiff {
			r++
		}
		farthest[i] = r
	}

	LOG := 1
	for (1 << LOG) < n {
		LOG++
	}

	up := make([][]int, LOG+1)
	up[0] = farthest
	for k := 1; k <= LOG; k++ {
		up[k] = make([]int, n)
		for i := 0; i < n; i++ {
			up[k][i] = up[k-1][up[k-1][i]]
		}
	}

    ans := make([]int, len(queries))
	for qi, q := range queries {
		u, v := pos[q[0]], pos[q[1]]
		if u > v {
			u, v = v, u
		}
		if u == v {
			ans[qi] = 0
			continue
		}
		cur, hops := u, 0
		for k := LOG; k >= 0; k-- {
			if up[k][cur] < v {
				hops += 1 << k
				cur = up[k][cur]
			}
		}
		if farthest[cur] >= v {
			ans[qi] = hops + 1
		} else {
			ans[qi] = -1
		}
	}
	return ans
}
