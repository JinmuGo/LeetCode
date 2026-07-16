func findMinArrowShots(points [][]int) int {
    slices.SortFunc(points, func(a, b []int) int {
        return cmp.Compare(a[1], b[1])
    })

    cnt := 1
    lastEnd := points[0][1]

    for i := 1; i < len(points); i++ {
        if points[i][0] > lastEnd {
            cnt++
            lastEnd = points[i][1]
        }
    }

    return cnt
}