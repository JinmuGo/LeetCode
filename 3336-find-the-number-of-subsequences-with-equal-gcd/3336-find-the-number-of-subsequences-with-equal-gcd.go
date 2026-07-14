func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func subsequencePairCount(nums []int) int {
    n := len(nums)
	mod := 1000000007
    maxVal := 200

    dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, maxVal+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, maxVal+1)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}

    getGCD := func(a, b int) int {
		if a == 0 {
			return b
		}
		if b == 0 {
			return a
		}
		return gcd(a, b)
	}

    var solve func(i, g1, g2 int) int
	solve = func(i, g1, g2 int) int {
		if i == n {
			if g1 > 0 && g2 > 0 && g1 == g2 {
				return 1
			}
			return 0
		}

		if dp[i][g1][g2] != -1 {
			return dp[i][g1][g2]
		}

		res := solve(i+1, g1, g2)

		res = (res + solve(i+1, getGCD(g1, nums[i]), g2)) % mod

		res = (res + solve(i+1, g1, getGCD(g2, nums[i]))) % mod

		dp[i][g1][g2] = res
		return res
	}

	return solve(0, 0, 0)
}