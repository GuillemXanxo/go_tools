package slice

// Min is the minimum value in the array and index where it is found, or zero.
func Min(ss []float64) (index int, min float64) {
	if len(ss) == 0 {
		return
	}

	min = ss[0]

	for i, s := range ss {
		if s < min {
			min = s
			index = i
		}
	}

	return
}
