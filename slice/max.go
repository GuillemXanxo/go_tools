package slice

// Max is the maximum value in the array and index where it is found, or zero.
func Max(ss []float64) (index int, max float64) {
	if len(ss) == 0 {
		return
	}

	max = ss[0]

	for i, s := range ss {
		if s > max {
			max = s
			index = i
		}
	}

	return
}
