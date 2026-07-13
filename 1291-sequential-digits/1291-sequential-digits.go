func sequentialDigits(low int, high int) []int {
    res := []int{}
    sample := "123456789"

    for length := 2; length <=9; length++ {
        for i := 0; i <= 9 - length; i++ {
            subStr := sample[i : i + length]
            num, _ := strconv.Atoi(subStr)

            if num >= low && num <= high {
                res = append(res, num)
            }
        }
    }   

    return res
}