func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []bool {
    prefix := make([]int, n)
    for i := 0; i < n-1; i++ {
        prefix[i+1] = prefix[i]
        if abs(nums[i] - nums[i+1]) > maxDiff {
            prefix[i+1]++
        }
    }
    
    res := make([]bool, len(queries))
	for i, query := range queries {
		u, v := min(query[0], query[1]), max(query[0], query[1])

        res[i] = (prefix[v] - prefix[u]) == 0
	}

    fmt.Println(prefix)

    return res
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}