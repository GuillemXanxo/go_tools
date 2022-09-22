package slice

//Filters all elements in a slice that satisfy the given predicate function
//We have to create a struct with a field that contains the function signature.
//Check explanaition:
func Filter[T any](ss []T, condition func(T) bool) (ss2 []T) {
	for _, s := range ss {
		if condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
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
	v := Filter(intSlice, try.f)
	fmt.Println(v) --> [1 3 4 2 3 1]

}
*/
