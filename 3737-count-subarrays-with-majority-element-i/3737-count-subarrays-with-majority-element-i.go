func countMajoritySubarrays(nums []int, target int) int {
    n := len(nums)
    pref := make([]int, n+1)
    ans := 0

    for i := 0; i < n; i++ {
        if nums[i] == target {
            pref[i + 1] = pref[i] + 1
        } else {
            pref[i + 1] = pref[i] - 1
        }
    }

    for l := 0; l < n; l++ {
        for r := l; r < n; r++ {
            if pref[r + 1] - pref[l] > 0 {
                ans++
            }
        }
    }
    
    return ans
}