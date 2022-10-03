package slice

// TakeBottom returns a new slice with `n` elements taken from the end.
func TakeBottom[S ~[]T, T any](slice S, n int) S {
	if n > len(slice) {
		n = len(slice)
	}
	return slice[len(slice)-n:]
}
