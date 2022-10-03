package slice

// Take returns a new slice with `n` elements taken from the beginning.
func TakeTop[S ~[]T, T any](slice S, n int) S {
	if n > len(slice) {
		n = len(slice)
	}
	return slice[:n]
}
