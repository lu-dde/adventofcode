package policy

import (
	"strconv"
	"strings"
)

//Pass struct for u2
type Pass struct {
	Min   int
	Max   int
	C     string
	Word  string
	Count int
}

//ParsePass split and return a Pass
func ParsePass(line string) Pass {
	a := strings.FieldsFunc(line, split)

	min, _ := strconv.Atoi(a[0])
	max, _ := strconv.Atoi(a[1])
	c := a[2]
	w := a[3]
	count := strings.Count(w, c)

	p := Pass{
		Min:   min,
		Max:   max,
		C:     c,
		Word:  w,
		Count: count,
	}

	return p
}

var splitMap = map[rune]bool{
	':': true,
	' ': true,
	'-': true,
}

func split(r rune) bool {
	return splitMap[r]
}
