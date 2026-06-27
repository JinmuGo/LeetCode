type Elem struct  {
    V string
    N int
}

func decodeString(s string) string {
    stack := make([]Elem, 0)
    var n int
    var ans string

    for _, c := range s {
        if '0' <= c && c <= '9' {
            n = n * 10 + int(c - '0')
        } else if c == '[' {
            stack = append(stack, Elem{V: ans, N: n})
            n = 0
            ans = ""
        } else if c == ']' {
            top := stack[len(stack) - 1]
            stack = stack[:len(stack) - 1]

            ans = top.V + strings.Repeat(ans, top.N)
        } else {
            ans += string(c)
        }
    }

    return ans
}