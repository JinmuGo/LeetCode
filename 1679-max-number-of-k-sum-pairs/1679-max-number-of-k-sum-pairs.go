func maxOperations(nums []int, k int) int {
    sort.Ints(nums)
    l, r := 0, len(nums) - 1
    ans := 0

    for l  < r {
        tmp := nums[l] + nums[r]

        if tmp == k {
            l++
            r--
            ans++
        } else if tmp > k {
            r--
        } else {
            l++
        }
    }

    return ans
}