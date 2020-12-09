package p1

import (
	"fmt"
	"strconv"
)

//Solve is main proxy for solve, takes a string channel
func Solve(p chan string, s chan string) {
	var t int64 = 0

	rulesLength := 25
	loop := make([]int64, rulesLength)
	for i := range loop {
		v, _ := strconv.ParseInt(<-p, 10, 64)
		loop[i] = v
	}

	pos := 0

	for line := range p {
		v, _ := strconv.ParseInt(line, 10, 64)

		//check
		if !sumOf2(v, loop) {
			t = v
			break
		}

		//save value and increment position
		loop[pos%rulesLength] = v
		pos++
	}

	s <- fmt.Sprintf("Solution: %d", t)
}

func sumOf2(sum int64, list []int64) bool {
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i]+list[j] == sum {
				return true
			}
		}
	}
	return false
}
