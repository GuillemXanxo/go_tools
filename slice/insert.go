package slice

// Insert a value at an index.
func Insert[T any](list []T, index int, values ...T) []T {
	if index >= len(list) {
		return append(list, values...)
	}

	return append(list[:index], append(values, list[index:]...)...)
}

/*
ingredients := ["banana", "cookie", "cheese"]
array = Insert(ingredients, 1, "chocolate")

print ingredients --> ["banana", "chocolate", "cookie", "cheese"]
*/
