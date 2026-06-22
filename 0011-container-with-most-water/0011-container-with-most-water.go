func maxArea(height []int) int {
    n := len(height)
    left, right := 0, n - 1
    ans := 0
    
    for left < right {
        minH := min(height[left], height[right])
        area := (right - left) * minH

        if area > ans {
            ans = area
        }

        if height[left] < height[right] {
            left++
        } else {
            right--
        }
    }

    return ans
}

