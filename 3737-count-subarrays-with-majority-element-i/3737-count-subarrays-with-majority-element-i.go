func countMajoritySubarrays(nums []int, target int) int {
    ans := 0
    n := len(nums)

    for l := 0; l < n; l++ {
        targetCnt := 0

        for r := l; r < n; r++ {
            if nums[r] == target {
                targetCnt++
            }

            len := r - l + 1

            if targetCnt > len / 2 {
                ans++
            }
        }
    }

    return ans
}