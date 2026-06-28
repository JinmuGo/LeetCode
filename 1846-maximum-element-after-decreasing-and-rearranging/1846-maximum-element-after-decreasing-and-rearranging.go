func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	sort.Ints(arr)
	curNum := 0

	for _, elem := range arr {
		if elem > curNum {
			curNum++
		}
	}

	return curNum
}