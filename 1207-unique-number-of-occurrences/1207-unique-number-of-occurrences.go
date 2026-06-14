func uniqueOccurrences(arr []int) bool {
    mapping := make(map[int]int)

    for _, num := range arr {
        _, ok := mapping[num]
        if ok {
            mapping[num] += 1
        } else {
            mapping[num] = 1
        }
    }

    same := make(map[int]bool)
    for _, val := range mapping {
        _, ok := same[val]

        if ok {
            return false
        } else {
            same[val] = true
        }
    }

    return true;
}