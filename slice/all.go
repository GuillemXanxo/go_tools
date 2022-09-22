package slice

// All will return true if all callbacks return true. It follows the same logic
// as the all() function in Python.
//
// If the list is empty then true is always returned.
func All[T any](ss []T, fn func(value T) bool) bool {
	for _, value := range ss {
		if !fn(value) {
			return false
		}
	}

	return true
}

/*
type Function struct {
	f func(int) bool
}

func callback(n int) bool {
	if n < 5 {
		return true
	} else {
		return false
	}
}

func main() {
	try := Function{f: callback}
	intSlice := []int{1, 5, 3, 6, 9, 9, 4, 2, 3, 1, 5}
	v := All(intSlice, try.f)
	fmt.Println(v) --> false

}
*/
