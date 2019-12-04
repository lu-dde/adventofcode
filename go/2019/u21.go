package main

import (
	"fmt"
	"strconv"
	"strings"
)

//U21 is main proxy for solve, takes a string channel
func U21(p chan string, s chan string) {

	// we only expect one line.
	line, _ := <-p

	code := strings.Split(line, ",")
	var ops = []int{}

	for _, c := range code {
		i, _ := strconv.Atoi(c)
		ops = append(ops, i)
	}

	ops[1] = 12
	ops[2] = 2

	step := 4
	var pos = 0

oploop:
	for {
		opcode := ops[pos]

		switch opcode {
		case 99:
			fmt.Println("99 :: break loop")
			break oploop
		case 1:
			source1 := ops[pos+1]
			source2 := ops[pos+2]
			values1 := ops[source1]
			values2 := ops[source2]
			target := ops[pos+3]
			ops[target] = values1 + values2
		case 2:
			source1 := ops[pos+1]
			source2 := ops[pos+2]
			values1 := ops[source1]
			values2 := ops[source2]
			target := ops[pos+3]
			ops[target] = values1 * values2
		}

		pos += step
	}

	s <- fmt.Sprintf("Solution: %d", ops[0])
}
