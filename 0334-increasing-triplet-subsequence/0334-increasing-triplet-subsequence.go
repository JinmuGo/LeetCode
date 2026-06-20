func increasingTriplet(nums []int) bool {
	first, second := math.MaxInt32, math.MaxInt32

	for _, num := range nums {
        switch {
            case num <= first:
                first = num
            case num <= second:
                second = num
            default:
                return true
        }
	}

	return false
}