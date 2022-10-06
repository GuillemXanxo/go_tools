package maps

//Checks if key exists in map m
func Has[K comparable, V any](m map[K]V, key K) bool {
	_, exists := m[key]
	return exists
}
