package maps

// Returns all keys present in both maps. Keys in second map are presented as "k-second"
// If wanted to work with other key types (now accepts only string keys):
// Change type of K in generics clause and adjust kPrima to match type.
func IntersectBothKeys[K string, V any](m1, m2 map[K]V) map[K]V {
	result := make(map[K]V)
	for k, v := range m1 {
		if Has(m2, k) {
			v2 := m2[k]
			kPrima := k + "-second"
			result[k] = v
			result[kPrima] = v2
		}
	}
	return result
}

// Intersect returns m1 values that keys are contained in m2.
// Only return k and v of m1
func Intersect[K comparable, V any](m1, m2 map[K]V) map[K]V {
	result := make(map[K]V)
	for k, v := range m1 {
		if Has(m2, k) {
			result[k] = v
		}
	}
	return result
}
