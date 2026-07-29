func smallestPalindrome(s string, k int) string {
	k64 := int64(k)
	var freq [26]int64
	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
	}

	var halfFreq [26]int64
	var halfLen int64 = 0
	mid := ""

	for i := 0; i < 26; i++ {
		halfFreq[i] = freq[i] / 2
		halfLen += halfFreq[i]
		if freq[i]%2 == 1 {
			mid = string(byte('a' + i))
		}
	}

	countPermutations := func(counts [26]int64, total int64) int64 {
		res := int64(1)
		currLen := int64(1)

		for i := 0; i < 26; i++ {
			c := counts[i]
			for j := int64(1); j <= c; j++ {
				res = res * currLen / j
				currLen++
				if res > k64 {
					return k64 + 1
				}
			}
		}
		return res
	}

	totalPerms := countPermutations(halfFreq, halfLen)
	if totalPerms < k64 {
		return ""
	}

	var leftBuilder strings.Builder

	for pos := 0; pos < int(halfLen); pos++ {
		for i := 0; i < 26; i++ {
			if halfFreq[i] == 0 {
				continue
			}

			halfFreq[i]--
			ways := countPermutations(halfFreq, halfLen-int64(pos)-1)

			if k64 <= ways {
				leftBuilder.WriteByte(byte('a' + i))
				break
			} else {
				k64 -= ways
				halfFreq[i]++
			}
		}
	}

	left := leftBuilder.String()

	runes := []rune(left)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	right := string(runes)

	return left + mid + right
}