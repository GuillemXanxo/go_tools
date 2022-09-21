package slice

//Returns a slice without repeated values
func uniqueIndexes[T comparable](ss []T) []T {
	keys := make(map[T]bool)
	list := []T{}
	for _, entry := range ss {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

/*
stringSlice := []string{"a", "b", "c", "d", "b", "a"}
uniqueString := uniqueIndexes(stringSlice)
print stringSlice --> ["a", "b", "c", "d", "b", "a"]
print uniqueString --> ["a", "b", "c", "d"]
*/
