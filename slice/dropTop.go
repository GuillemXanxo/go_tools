package slice

// DropTop will return the rest slice after dropping the top n elements
// if the slice has less elements then n that'll return empty slice
// if n < 0 it'll return empty slice.
func DropTop[T any](ss []T, n int) (drop []T) {
	if n < 0 || n >= len(ss) {
		return
	}
	drop = make([]T, len(ss)-n)
	copy(drop, ss[n:])

	return
}
