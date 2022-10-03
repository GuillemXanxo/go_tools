package slice

// Partition creates two slices, the first of which contains elements that
// `predicate` returns true for, with the second containing elements for which
// `predicate` returns false.
func Partition[S ~[]T, T any](slice S, predicate func(T) bool) (truths S, falsehoods S) {
	truths = make(S, 0)
	falsehoods = make(S, 0)
	for _, item := range slice {
		if predicate(item) {
			truths = append(truths, item)
		} else {
			falsehoods = append(falsehoods, item)
		}
	}
	return
}

/*
type Function struct {
	f func(int) bool
}

func callback(n int) bool {
	if n > 5 {
		return true
	} else {
		return false
	}
}

func main() {
	try := Function{f: callback}
	intSlice := []int{1, 5, 3, 6, 9, 9, 4, 2, 3, 1, 5}
	fmt.Println(intSlice)
	truths, falses := Partition(intSlice, try.f)
	fmt.Println(truths) --> [6, 9, 9]
	fmt.Println(falses) --> [1, 5, 3, 4, 2, 3, 1, 5]
}
*/
