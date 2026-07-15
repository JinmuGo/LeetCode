func gcdOfOddEvenSums(n int) int {
	sumOdd := n * n
	sumEven := n * (n + 1)

    return gcd(sumOdd, sumEven)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}