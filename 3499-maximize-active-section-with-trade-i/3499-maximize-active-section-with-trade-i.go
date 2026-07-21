type ZeroOneSegment struct {
	zeroCount int
	oneCount  int
}

func maxActiveSectionsAfterTrade(s string) int {
	n := len(s)
	totalOnes := 0

	for i := 0; i < n; i++ {
		if s[i] == '1' {
			totalOnes++
		}
	}

	i := 0
	for i < n && s[i] == '1' {
		i++
	}

	var segments []ZeroOneSegment

	for i < n {
		zeros := 0
		for i < n && s[i] == '0' {
			zeros++
			i++
		}

		ones := 0
		for i < n && s[i] == '1' {
			ones++
			i++
		}

		segments = append(segments, ZeroOneSegment{
			zeroCount: zeros,
			oneCount:  ones,
		})
	}

	if len(segments) < 2 {
		return totalOnes
	}

	maxDelta := 0
	for idx := 0; idx < len(segments)-1; idx++ {
		if segments[idx].oneCount > 0 {
			delta := segments[idx].zeroCount + segments[idx+1].zeroCount
			if delta > maxDelta {
				maxDelta = delta
			}
		}
	}

	return totalOnes + maxDelta
}