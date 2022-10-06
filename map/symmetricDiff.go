package maps

//Returns a map with all keys that don't exist in both maps
func SymmetricDiff[K comparable, V any](m1, m2 map[K]V) map[K]V {
	result := Clone(m1)
	for k, v := range m2 {
		if Has(m1, k) {
			delete(result, k)
		} else {
			result[k] = v
		}
	}
	return result
}
