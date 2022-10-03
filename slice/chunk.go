package slice

import "math"

// Chunk splits the given slice into smaller slices, each the length of `chunkSize`.
// If the slice cannot be split evenly, the last chunk will have the remaining elements.
func Chunk[S ~[]T, T any](slice S, chunkSize int) []S {
	chunks := int(math.Ceil(float64(len(slice)) / float64(chunkSize)))
	output := make([]S, chunks)
	for c := 0; c < chunks; c++ {
		start := c * chunkSize
		end := start + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		output[c] = slice[start:end]
	}
	return output
}
