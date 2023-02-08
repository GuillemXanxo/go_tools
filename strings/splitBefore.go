package strings

import (
	"errors"
	"strings"
)

//Splits string in 2 parts before the sep
/*
string := email@dot.com
separator := "@"
part1 --> email
part2 --> @dot.com
*/

func SplitBefore(s, separator string) (part1, part2 string, e error) {
	if i := strings.Index(s, separator); i >= 0 {
		part1, part2 := s[:i], s[i:]
		return part1, part2, nil
	} else {
		e := errors.New("Cannot find separator in s")
		return "", "", e
	}
}
