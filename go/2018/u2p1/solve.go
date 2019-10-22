package main

import (
	"fmt"
	"log"
)

func solve(p chan string, s chan string) {

	sum := 0

	for {
		line, ok := <-p
		if ok {
			log.Print(line)
		} else {
			break
		}
	}

	s <- fmt.Sprintf("Solution: %d", sum)
	return
}
