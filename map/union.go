package maps

//Returns a map that is m1 + m2 where in case of same keys, values from m1 are preferred
func Union[K comparable, V any](m1, m2 map[K]V) map[K]V {
	result := make(map[K]V)
	for k, v := range m2 {
		result[k] = v
	}
	for k, v := range m1 {
		result[k] = v
	}
	return result
}
