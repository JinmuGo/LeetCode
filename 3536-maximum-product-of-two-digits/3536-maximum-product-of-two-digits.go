func maxProduct(n int) int {
    str := []rune(strconv.Itoa(n))

    sort.Slice(str, func(i, j int) bool {
        return str[i] > str[j]
    })

    return int(str[0] - '0') * int(str[1] - '0')
}