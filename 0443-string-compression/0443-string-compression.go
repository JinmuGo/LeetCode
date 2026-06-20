func compress(chars []byte) int {
    i := 0
    idx := 0
    n := len(chars)

    for i < n {
        cur := chars[i]
        cnt := 0

        for i < n && cur == chars[i] {
            i++
            cnt++
        }

        chars[idx] = cur
        idx++

        if cnt > 1 {
            for _, c := range strconv.Itoa(cnt) {
                chars[idx] = byte(c)
                idx++
            }   
        }
    }

    return idx
}