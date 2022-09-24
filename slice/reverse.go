package slice

// Reverse returns a new copy of the slice with the elements ordered in reverse.
func Reverse[T any](ss []T) []T {
	if len(ss) < 2 {
		return ss
	}

	sorted := make([]T, len(ss))
	for i := 0; i < len(ss); i++ {
		sorted[i] = ss[len(ss)-i-1]
	}

	return sorted
}
