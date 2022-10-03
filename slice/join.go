package slice

import (
	"fmt"
	"strings"
)

// Join concatenates all the elements of a slice into a string separated by `separator`.
// `fmt.Sprint` is used for to get the string representation of the given value, so mixed types
// are possible with `[]any`.
func Join[S ~[]T, T any](slice S, separator string) string {
	stringList := make([]string, len(slice))
	for i, e := range slice {
		stringList[i] = fmt.Sprint(e)
	}
	return strings.Join(stringList, separator)
}
