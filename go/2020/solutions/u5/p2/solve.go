package p2

import (
	"fmt"
)

//Solve is main proxy for solve, takes a string channel
func Solve(p chan string, s chan string) {
	var t = 0

	for line := range p {
		row := 0
		col := 0
		for i := 0; i < 7; i++ {
			row <<= 1
			if line[i] == 'B' {
				row++
			}
		}
		for i := 0; i < 3; i++ {
			col <<= 1
			if line[i+7] == 'R' {
				col++
			}
		}

		v := row*8 + col
		if t < v {
			t = v
		}
	}

	s <- fmt.Sprintf("Solution: %d", t)
}
