func maximumProduct(nums []int) int {
    sort.Ints(nums)
    n := len(nums)

    c1 := nums[n - 1] * nums[n - 2] * nums[n - 3]
    c2 := nums[0] * nums[1] * nums [n - 1]

    return max(c1, c2)
}