func uniqueOccurrences(arr []int) bool {
    mapping := make(map[int]int)

    for _, num := range arr {
        mapping[num]++
    }

    same := make(map[int]struct{})
    for _, val := range mapping {
        if _, ok := same[val]; ok {
            return false
        } else {
            same[val] = struct{}{}
        }
    }

    return true;
}