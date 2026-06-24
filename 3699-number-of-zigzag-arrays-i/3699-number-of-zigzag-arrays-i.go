const (
	MOD      = 1000000007 
	DOWN, UP = 0, 1      
)

func zigZagArrays(n int, l int, r int) int {
	dp := make([][][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][]int, 2)
		for j := 0; j < 2; j++ {
			dp[i][j] = make([]int, r+1)
		}
	}

	for x := l; x <= r; x++ {
		dp[1][DOWN][x] = 1
		dp[1][UP][x] = 1
	}

	for i := 2; i <= n; i++ {
		prefUp := make([]int, r+2)
		sumUp := 0
		for y := l; y <= r; y++ {
			prefUp[y] = sumUp
			sumUp = (sumUp + dp[i-1][UP][y]) % MOD
		}

		prefDown := make([]int, r+2)
		sumDown := 0
		for y := r; y >= l; y-- {
			prefDown[y] = sumDown
			sumDown = (sumDown + dp[i-1][DOWN][y]) % MOD
		}

		for x := l; x <= r; x++ {
			dp[i][DOWN][x] = prefUp[x]
			dp[i][UP][x] = prefDown[x]
		}
	}

	ans := 0
	for x := l; x <= r; x++ {
		ans = (ans + dp[n][DOWN][x] + dp[n][UP][x]) % MOD
	}

	return ans
}