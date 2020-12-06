package p2

import (
	"fmt"
)

//Solve is main proxy for solve, takes a string channel
func Solve(p chan string, s chan string) {
	var t = 0

	counts := []map[byte]int{}

	group := 0
	for line := range p {
		if line == "" {
			group++
			continue
		}

		if len(counts) == group {
			counts = append(counts, map[byte]int{})
		}

		counts[group][0]++

		for i := 0; i < len(line); i++ {
			counts[group][line[i]]++
		}
	}

	for _, c := range counts {
		gt := -1
		nr := c[0]
		for _, answers := range c {
			if nr == answers {
				gt++
			}
		}
		//fmt.Println("+", gt)

		t += gt
	}

	s <- fmt.Sprintf("Solution: %d", t)
}
