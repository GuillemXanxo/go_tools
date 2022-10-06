package maps

//Returns all the keys that are in m2 that are not present in m1
func Diff[K comparable, V1, V2 any](m1 map[K]V1, m2 map[K]V2) map[K]V1 {
	result := make(map[K]V1)
	for k, v := range m1 {
		if !Has(m2, k) {
			result[k] = v
		}
	}
	return result
}
