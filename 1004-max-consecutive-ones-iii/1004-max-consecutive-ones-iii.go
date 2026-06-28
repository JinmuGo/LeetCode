func longestOnes(nums []int, k int) int {
    n := len(nums)
    left, zero, maxVal := 0, 0, 0

    for right := 0; right < n; right++ {
        if nums[right] == 0 {
            zero++
        }
        for zero > k {
            if nums[left] == 0 {    
                zero--
            }
            left++
        }

        maxVal = max(maxVal, right - left + 1)

    }

    return maxVal
}