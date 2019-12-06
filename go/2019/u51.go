package main

import (
	"fmt"
	"strconv"
	"strings"
)

//U51 is main proxy for solve, takes a string channel
func U51(p chan string, s chan string) {

	// we only expect one line.
	line, _ := <-p

	code := strings.Split(line, ",")
	var cps = []int{}

	for _, c := range code {
		i, _ := strconv.Atoi(c)
		cps = append(cps, i)
	}

	opcodeChan := make(chan []int)
	solutions := make(chan sol)

	go opscode2(opcodeChan, solutions)

	opcodeChan <- cps
	close(opcodeChan)

	solution, _ := <-solutions
	s <- fmt.Sprintf("Solution: %d, %v", 100*solution.noun+solution.verb, solution)

}

type sol2 struct {
	noun   int
	verb   int
	result int
}

func opscode2(opsChan chan []int, solution chan sol) {

	for ops := range opsChan {
		n := ops[1]
		v := ops[2]

		var pos = 0

	oploop:
		for {
			opcode := ops[pos]

			switch opcode {
			case 99:
				break oploop
			case 1:
				source1 := ops[pos+1]
				source2 := ops[pos+2]
				values1 := ops[source1]
				values2 := ops[source2]
				target := ops[pos+3]
				ops[target] = values1 + values2
				pos += 4
			case 2:
				source1 := ops[pos+1]
				source2 := ops[pos+2]
				values1 := ops[source1]
				values2 := ops[source2]
				target := ops[pos+3]
				ops[target] = values1 * values2
				pos += 4
			case 3:
				source1 := ops[pos+1]
				ops[source1] = getIntCodeInput()
				pos += 2
			case 4:
				source1 := ops[pos+1]
				putIntCodeOutput(ops[source1])
				pos += 2
			}

		}

		solution <- sol{noun: n, verb: v, result: ops[0]}
	}
}

func getIntCodeInput() int {
	fmt.Printf("read: %d\n", 0)
	return 31
}

func putIntCodeOutput(output int) {
	fmt.Printf("put: %d\n", output)
}
