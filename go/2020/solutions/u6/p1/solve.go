package p1

import (
	"bytes"
	"fmt"
)

//Solve is main proxy for solve, takes a string channel
func Solve(p chan string, s chan string) {
	var t = 0

	counts := [][]byte{}

	group := 0
	for line := range p {
		if line == "" {
			group++
			continue
		}

		if len(counts) == group {
			counts = append(counts, []byte{})
		}

		for i := 0; i < len(line); i++ {
			if bytes.IndexByte(counts[group], line[i]) == -1 {
				counts[group] = append(counts[group], line[i])
			}
		}
	}

	//fmt.Println(counts)

	for _, c := range counts {
		t += len(c)
	}

	s <- fmt.Sprintf("Solution: %d", t)
}
