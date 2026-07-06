func removeCoveredIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int) bool {
        if intervals[i][0] == intervals[j][0] {
            return intervals[i][1] > intervals[j][1]
        }
        return intervals[i][0] < intervals[j][0]
    })

    cnt := 0
    prevEnd := 0

    for _, interval := range intervals {
        end := interval[1]

        if end > prevEnd {
            cnt++
            prevEnd = end
        }
    }

    return cnt
}