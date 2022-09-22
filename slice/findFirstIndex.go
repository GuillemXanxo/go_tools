package slice

//Returns first index that satisfy condition set by function in parameter.
//If no index satisfies, then returns -1.
func FindFirstIndex[T any](ss []T, fn func(T) bool) int {
	for idx, value := range ss {
		if fn(value) {
			return idx
		}
	}

	return -1
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
	v := FindFirstIndex(intSlice, try.f)
	fmt.Println(v) --> 3

}
*/
