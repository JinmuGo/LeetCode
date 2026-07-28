func smallestPalindrome(s string) string {
    var alpha [26]int
    
    for _, c := range s {
        alpha[c - 'a']++
    }

    var leftBuilder strings.Builder
    mid := ""

    for i := 0; i < 26; i++ {
        char := byte('a' + i)
        count := alpha[i]

        if count > 0 {
            leftBuilder.WriteString(strings.Repeat(string(char), count / 2))

            if count % 2 == 1 {
                mid = string(char)
            }
        }
    }

    left := leftBuilder.String()

    runes := []rune(left)
    for i, j := 0, len(runes) - 1; i < j; i, j = i + 1, j - 1 {
        runes[i], runes[j] = runes[j], runes[i]
    }

    right := string(runes)

    return left + mid + right
}