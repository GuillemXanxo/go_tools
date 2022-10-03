package slice

// Concat combines all the elements from all the given slices into a single slice.
// Still have to be slices of the same type.
func Concat[S ~[]T, T any](slices ...S) S {
	output := make(S, 0)
	for _, list := range slices {
		output = append(output, list...)
	}
	return output
}
