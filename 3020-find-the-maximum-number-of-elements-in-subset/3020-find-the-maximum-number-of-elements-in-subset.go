func maximumLength(nums []int) int {
    maps := make(map[int]int, 0)

    for _, num := range nums {
        maps[num]++
    }

    maxVal := 1

    if ones, ok := maps[1]; ok {
        if ones%2 == 0 {
            maxVal = max(maxVal, ones-1)
        } else {
            maxVal = max(maxVal, ones)
        }
    }

    dp := make(map[int]int)

    for num := range maps {
        if num == 1 {
            continue
        }

        if _, ok := dp[num]; ok {
            continue
        }

        cur := num
        var chain []int

        for maps[cur] >= 2 {
            chain = append(chain, cur)
            cur = cur * cur
        }

        totalLen := 0
        if maps[cur] >= 1 {
            totalLen++
        } else  {
            totalLen--
        }

        for i := len(chain) - 1; i >= 0; i-- {
            totalLen += 2
            dp[chain[i]] = totalLen
            maxVal = max(maxVal, totalLen)
        } 
    }

    return maxVal
}