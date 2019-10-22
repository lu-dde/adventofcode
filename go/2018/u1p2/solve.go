package main

import (
	"fmt"
	"strconv"
)

func solve(p chan string, s chan string) {

	values := [1018]int{}
	index := 0

	for {
		line, ok := <-p
		if ok {
			sign := line[0]
			rest := line[1:]
			i, err := strconv.Atoi(rest)
			if err == nil {
				if sign == '+' {
					values[index] = i
				} else {
					values[index] = -i
				}
			}
		} else {
			break
		}
		index++
	}

	sum := 0
	passed := map[int]bool{0: true}

	for {
		for _, value := range values {
			sum += value

			_, exists := passed[sum]
			if exists {
				s <- fmt.Sprintf("Solution: %d", sum)
				return
			}
			passed[sum] = true
		}
	}
}
