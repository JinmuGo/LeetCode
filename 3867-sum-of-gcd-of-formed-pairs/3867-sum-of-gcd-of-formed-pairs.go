func gcdSum(nums []int) int64 {
    n := len(nums)
    mx := 0
    prefixGcd := make([]int, n)

    for i, num := range nums {
        if num > mx {
            mx = num
        }
        prefixGcd[i] = gcd(num, mx)
    }

    sort.Ints(prefixGcd)
    
    l, r := 0, n - 1
    var s int64

    for l < r {
        lE, rE := prefixGcd[l], prefixGcd[r]
        s += int64(gcd(lE, rE))
        l++
        r--
    }

    return s
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}