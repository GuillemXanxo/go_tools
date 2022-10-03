package slice

// Remove returns a new slice without the elements for which the `predicate`
// returns `true`.
func Remove[S ~[]T, T any](slice S, predicate func(value T, index int, slice S) bool) S {
	output := make(S, 0)
	for i := range slice {
		if !predicate(slice[i], i, slice) {
			output = append(output, slice[i])
		}
	}
	return output
}

/*
func main() {
	try := Function{f: No_B}
	uniqueSlice := []string{"a", "b", "c", "d", "e", "f"}
	v := Remove(uniqueSlice, try.f)
	fmt.Println(v) --> [a, c, d, e, f]
}

type Function struct {
	f func(string, int, []string) bool
}

func No_B(a string, b int, slice []string) bool {
	if a == "b" {
		return true
	} else {
		return false
	}
}
*/
