func gcdValues(nums []int, queries []int64) []int {
    maxVal := slices.Max(nums)

    freq := make([]int, maxVal + 1)
    for _, num := range nums {
        freq[num]++
    }

    exactPairs := make([]int64, maxVal + 1)

    for i := maxVal; i >= 1; i-- {
        var cnt int64

        for j := i; j <= maxVal; j += i {
            cnt += int64(freq[j])
        }

        pairs := cnt * (cnt - 1) / 2

        for j := 2 * i; j <= maxVal; j += i {
            pairs -= exactPairs[j]
        }

        exactPairs[i] = pairs
    }

    pref := make([]int64, maxVal + 1)
    for i := 1; i <= maxVal; i++ {
        pref[i] = pref[i - 1] + exactPairs[i]
    }

    ans := make([]int, len(queries))
    for i, q := range queries {
        g := sort.Search(len(pref), func(x int) bool {
            return pref[x] > q
        })
        ans[i] = g
    }

    return ans
}
