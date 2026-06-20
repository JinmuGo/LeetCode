func compress(chars []byte) int {
	cnt, idx := 1, 0
	n := len(chars)

	for i := 1; i <= n; i++ {
		if i < n && chars[i] == chars[i - 1] {
			cnt++
		} else {
			chars[idx] = chars[i - 1]
			idx++

			if cnt > 1 {
				for _, c := range strconv.Itoa(cnt) {
					chars[idx] = byte(c)
					idx++
				}
			}
			cnt = 1
		}
	}

	return idx
}