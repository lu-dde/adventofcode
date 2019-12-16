package main

import (
	"fmt"
	"strconv"
)

//U101 is main proxy for solve, takes a string channel
func U101(p chan string, s chan string) {
	var t = 0

	for {
		line, ok := <-p
		if ok {
			i, _ := strconv.Atoi(line)
			t += i/3 - 2
		} else {
			break
		}
	}

	s <- fmt.Sprintf("Solution: %d", t)
}
