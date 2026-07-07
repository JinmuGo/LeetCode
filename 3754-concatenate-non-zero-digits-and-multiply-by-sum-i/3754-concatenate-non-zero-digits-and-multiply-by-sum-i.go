func sumAndMultiply(n int) int64 {
    str := strconv.Itoa(n)
    x := ""
    var sum int64

    for _, s := range str {
        if s == '0' {
            continue
        }
        x += string(s)
        sum += int64(s - '0')
    }

    num, _ := strconv.Atoi(x)

    sum *= int64(num)

    return sum
}