package main

import (
	"fmt"
)

func counts(l string) (int, int) {
	m := make(map[rune]int32, 26)
	for _, c := range l {
		_, prev := m[c]
		if prev {
			m[c]++
		} else {
			m[c] = 1
		}
	}

	c2 := 0
	c3 := 0

	for _, value := range m {
		if value == 2 {
			c2 = 1
		}
		if value == 3 {
			c3 = 1
		}
	}

	return c2, c3
}

func solve(p chan string, s chan string) {

	c2 := 0
	c3 := 0

	for {
		line, ok := <-p
		if ok {
			nc2, nc3 := counts(line)
			c2 += nc2
			c3 += nc3
		} else {
			break
		}
	}

	s <- fmt.Sprintf("Solution: %d", c2*c3)
	return
}
