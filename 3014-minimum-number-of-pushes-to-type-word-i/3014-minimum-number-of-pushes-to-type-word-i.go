func minimumPushes(word string) int {
	n := len(word)
	v, c := n%8, n/8

	return 4*c*(c+1) + v*(c+1)
}