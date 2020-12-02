package p2

import (
	"fmt"
	"strconv"
	"strings"
)

//Solve is main proxy for solve, takes a string channel
func Solve(p chan string, s chan string) {
	var t = 0

	for line := range p {
		if testPass(line) {
			t++
		}
	}

	s <- fmt.Sprintf("Solution: %d", t)
}

//ParsePass2 split and return a Pass
func testPass(line string) bool {
	a := strings.FieldsFunc(line, split)

	p1, _ := strconv.Atoi(a[0])
	p2, _ := strconv.Atoi(a[1])
	c := a[2][0]
	w := a[3]

	return (w[p1-1] == c) != (w[p2-1] == c)
}

var splitMap = map[rune]bool{
	':': true,
	' ': true,
	'-': true,
}

func split(r rune) bool {
	return splitMap[r]
}
