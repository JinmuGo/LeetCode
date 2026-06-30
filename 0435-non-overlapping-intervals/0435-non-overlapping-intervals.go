func eraseOverlapIntervals(intervals [][]int) int {
    // minium number
    // 구간의 길이가 긴걸 지우면 좋음. 

    if len(intervals) <= 1 {
        return 0
    }

    slices.SortFunc(intervals, func(a, b []int) int {
        return a[1] - b[1]
    })

    cnt := 0
    prevEnd := intervals[0][1]
    
    for i := 1; i < len(intervals); i++ {
        start := intervals[i][0]
        end := intervals[i][1]

        if start < prevEnd {
            cnt++
        } else {
            prevEnd = end
        }
    }

    return cnt
}