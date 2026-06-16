func pivotIndex(nums []int) int {
    left, right := 0, 0

    for _, num := range nums {
        right += num
    }

    for j, val := range nums {
        right -= val
        if left == right {
            return j
        }
        left += val
    }

    return -1
}