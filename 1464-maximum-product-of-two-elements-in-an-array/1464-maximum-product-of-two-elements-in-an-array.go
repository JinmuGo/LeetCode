func maxProduct(nums []int) int {
    big, sbig := 0, 0

    for _, num := range nums {
        if num > big {
            sbig = big
            big = num
        } else {
            sbig = max(sbig, num)
        }
    }

    return (big - 1) * (sbig - 1)
}