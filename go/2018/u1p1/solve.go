package main

import (
	"fmt"
	"log"
	"strconv"
)

func solve(p chan string, s chan string) {

	sum := 0

	for {
		line, ok := <-p
		if ok {
			sign := line[0]
			rest := line[1:]
			i, err := strconv.Atoi(rest)
			if err == nil {
				if sign == '+' {
					sum += i
				} else {
					sum -= i
				}
			} else {
				log.Printf("failed to parse <%s>", line)
			}
		} else {
			break
		}
	}

	s <- fmt.Sprintf("Solution: %d", sum)
	return
}
