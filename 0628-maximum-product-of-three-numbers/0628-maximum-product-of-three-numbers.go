func maximumProduct(nums []int) int {
    sort.Slice(nums, func(i, j int) bool {
        return nums[i] > nums[j]
    })

    n := len(nums)

    candidate := nums[0] * nums[1] * nums[2]
    negative := nums[0] * nums[n - 1] * nums [n - 2]

    if negative > candidate {
        return negative
    }
    return candidate
}