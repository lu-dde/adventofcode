package p2

import (
	"fmt"
)

//Solve is main proxy for solve, takes a string channel
func Solve(input chan string, s chan string) {
	var t int64 = 0

	var r1 int64 = 0
	var r2 int64 = 0
	var r3 int64 = 0
	var r5 int64 = 0
	var r7 int64 = 0

	var rv1 int64 = 0
	var rv2 int64 = 0
	var rv3 int64 = 0
	var rv5 int64 = 0
	var rv7 int64 = 0

	l := 0

	for line := range input {
		width := int64(len(line))
		//fmt.Println(l, line, r1, r3, r5, r7, r2, "\t=", rv1, "*", rv3, "*", rv5, "*", rv7, "*", rv2)

		if line[r1] == '#' {
			rv1++
		}

		r1++
		r1 = r1 % width

		if line[r3] == '#' {
			rv3++
		}

		r3 += 3
		r3 = r3 % width

		if line[r5] == '#' {
			rv5++
		}

		r5 += 5
		r5 = r5 % width

		if line[r7] == '#' {
			rv7++
		}

		r7 += 7
		r7 = r7 % width

		if line[r2] == '#' {
			if l%2 == 0 {
				rv2++
			}
		}

		if l%2 == 0 {
			r2++
		}
		r2 = r2 % width
		//fmt.Println(l, line, r1, r3, r5, r7, r2, "\t=", rv1, "*", rv3, "*", rv5, "*", rv7, "*", rv2)

		l++
	}

	//fmt.Println(rv1, "*", rv3, "*", rv5, "*", rv7, "*", rv2)
	t = rv1 * rv3 * rv5 * rv7 * rv2

	s <- fmt.Sprintf("Solution: %d", t)
}
