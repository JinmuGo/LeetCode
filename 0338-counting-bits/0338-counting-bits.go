func countBits(n int) []int {
    ans := make([]int, n + 1)

    for i := 0; i <= n; i++ {
        ans[i] = popCount(i)
    }

    return ans
}

func popCount(x int) int {
    count := 0

    for x > 0 {
        count += x & 1
        x >>= 1
    }

    return count
}