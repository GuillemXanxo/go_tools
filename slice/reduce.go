package slice

// Reduce continually applies the provided function
// over the slice. Reducing the elements to a single value.
func Reduce(ss []float64, reducer func(float64, float64) float64) (el float64) {
	if len(ss) == 0 {
		return
	}
	el = ss[0]
	for _, s := range ss[1:] {
		el = reducer(el, s)
	}
	return
}

/*
func reducer(a, b float64) float64 {
	return a + b
}

type Function struct {
	f func(float64, float64) float64
}

func main() {
	try := Function{f: reducer}
	intSlice := []float64{1, 5, 3, 6, 9, 9, 4, 2, 3, 1, 5}
	v := Reduce(intSlice, try.f)
	fmt.Println(v) --> 1+5+3+6+9+9+4+2+3+1+5=48
}
*/
