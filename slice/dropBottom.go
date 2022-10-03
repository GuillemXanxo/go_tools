package slice

// DropBottom returns a new slice with `n` elements dropped from the end.
func DropBottom[S ~[]T, T any](slice S, n int) S {
	if n > len(slice) {
		n = len(slice)
	}
	return slice[:len(slice)-n]
}
