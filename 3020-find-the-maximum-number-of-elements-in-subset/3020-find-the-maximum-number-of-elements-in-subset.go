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

    var recursive func(num, cnt, res int) int

    recursive = func (num, cnt, res int) int {
        if num <= 1 {
            return res
        }
        if cnt2, ok := maps[num]; !ok || cnt2 < cnt {
            return res
        }

        res += cnt

        nextNum := int(math.Sqrt(float64(num)))

        if nextNum * nextNum != num {
            return res
        }

        return recursive(nextNum, 2, res)

    }

    for num := range maps {
        if num == 1 {
            continue
        }
        maxVal = max(maxVal, recursive(num, 0, 1))
    }

    return maxVal
}