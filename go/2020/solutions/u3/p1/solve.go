package p1

import (
	"fmt"
)

//Solve is main proxy for solve, takes a string channel
func Solve(input chan string, s chan string) {
	var t = 0

	p := 0

	for line := range input {

		//fmt.Println(p, line)

		if line[p] == 35 {
			t++
		}

		p += 3
		p = p % len(line)
	}

	s <- fmt.Sprintf("Solution: %d", t)
}
