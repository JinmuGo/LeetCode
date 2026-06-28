func dailyTemperatures(temperatures []int) []int {
    ans := make([]int, len(temperatures))
    stack := make([]int, 0)

    for i, temp := range temperatures {
        for len(stack) > 0 && temp > temperatures[stack[len(stack) - 1]] {
            idx := stack[len(stack) - 1]
            stack = stack[:len(stack) - 1]

            ans[idx] = i - idx
        }

        stack = append(stack, i)
    }
    
    return ans
}