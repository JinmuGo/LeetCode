func maxProduct(n int) int {
    first, second := 0, 0

	for n > 0 {
		x := n % 10
		if x > first {
			second = first
			first = x
		} else if x > second {
			second = x
		}
		n /= 10
	}
    
	return first * second
}