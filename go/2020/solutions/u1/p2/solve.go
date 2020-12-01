package p2

import (
	"fmt"
	"strconv"
)

//Solve2 is main proxy for solve, takes a string channel
func Solve2(p chan string, s chan string) {
	var t = 0

	var want map[int]int = make(map[int]int, 200)

	for line := range p {
		i, _ := strconv.Atoi(line)
		want[2020-i] = i
	}

outer:
	for _, w := range want {
		for _, e := range want {
			if want[w+e] > 0 {
				f := want[w+e]
				//fmt.Println("found", w, "+", e, "+", f, "=", w+e+f)
				t = w * e * f
				break outer
			}
		}
	}

	s <- fmt.Sprintf("Solution: %d", t)
}
