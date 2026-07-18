func findGCD(nums []int) int {
    min := 1001
    max := 0

    for _, num := range nums {
        if num < min {
            min = num
        }
        if num > max {
            max = num
        }
    }

    return gcd(min, max)
}

func gcd (a, b int) int {
    for b != 0 {
        a, b = b, a % b
    }

    return a
}