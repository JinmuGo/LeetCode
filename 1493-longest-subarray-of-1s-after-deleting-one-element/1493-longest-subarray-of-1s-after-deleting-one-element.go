func longestSubarray(nums []int) int {
    left := 0
    zeroCnt := 0
    maxLen := 0

    for right, elem := range nums {
        if elem == 0 {
            zeroCnt++
        }

        for zeroCnt > 1 {
            if nums[left] == 0 {
                zeroCnt--
            }
            left++
        }
        maxLen = max(maxLen, right - left)
    }
    
    return maxLen
}