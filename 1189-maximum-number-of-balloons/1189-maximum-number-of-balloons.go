func maxNumberOfBalloons(text string) int {
	table := map[rune]int{
		'b': 0,
		'a': 0,
		'l': 0,
		'o': 0,
		'n': 0,
	}

	for _, c := range text {
		if _, ok := table[c]; ok {
			table[c]++
		}
	}

	return min(table['b'], table['a'], table['l']/2, table['o']/2, table['n'])
}