func arrayRankTransform(arr []int) []int {
    n := len(arr)
    duplicate := make([]int, n)
    copy(duplicate, arr)

    sort.Ints(duplicate)

    ranks := make(map[int]int)
    rank := 1

    for _, v := range duplicate {
        if _, exists := ranks[v]; !exists {
            ranks[v] = rank
            rank++
        }
    }

    res := make([]int, n)
    for i, v := range arr {
        res[i] = ranks[v]
    }

    return res
}