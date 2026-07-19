func smallestSubsequence(s string) string {
    count := make(map[byte]int)
    for i := 0; i < len(s); i++ {
        count[s[i]]++
    }

    visited := make(map[byte]bool)
    stack := []byte{}

    for i := 0; i < len(s); i++ {
        char := s[i]
        count[char]-- 

        if visited[char] {
            continue
        }

        for len(stack) > 0 && stack[len(stack)-1] > char && count[stack[len(stack)-1]] > 0 {
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            visited[top] = false       
        }
        
        stack = append(stack, char)
        visited[char] = true
    }

    return string(stack)
}