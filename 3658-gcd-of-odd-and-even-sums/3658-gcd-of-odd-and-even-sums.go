func gcdOfOddEvenSums(n int) int {
	odd := 1
	even := 2
	sumOdd, sumEven := 0, 0

	for i := 0; i < n; i++ {
		sumOdd += odd
		sumEven += even

		odd += 2
		even += 2
	}


    return gcd(sumOdd, sumEven)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}