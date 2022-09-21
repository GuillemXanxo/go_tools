package main

import (
	"fmt"
)

func Contains[T comparable](ss []T, lookingFor T) bool {
	for _, s := range ss {
		if lookingFor == s {
			return true
		}
	}

	return false
}

func main() {
	intSlice := []int{1, 5, 3, 6, 9, 9, 4, 2, 3, 1, 5}
	fmt.Println(intSlice)
	uniqueSlice := Contains(intSlice, 3)
	fmt.Println(uniqueSlice)

}
