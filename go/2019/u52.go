package main

import (
	"fmt"
	"strconv"
	"strings"
)

//U52 is main proxy for solve, takes a string channel
func U52(p chan string, s chan string) {

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

	go opscode3(opcodeChan, solutions)

	opcodeChan <- cps
	close(opcodeChan)

	solution, _ := <-solutions
	s <- fmt.Sprintf("Solution: %d, %v", 100*solution.noun+solution.verb, solution)

}

func opscode3(opsChan chan []int, solution chan sol) {

	for ops := range opsChan {
		n := ops[1]
		v := ops[2]

		var pos = 0

	oploop:
		for {
			opcodeCompact := ops[pos]

			opcode := opcodeCompact % 100
			p1Mode := (opcodeCompact / 100) % 10
			p2Mode := (opcodeCompact / 1000) % 10
			//p3Mode := (opcodeCompact / 10000) % 10

			//fmt.Println("opcode", opcodeCompact, opcode, p1Mode, p2Mode, p3Mode)

			var pos1 = pos + 1
			var pos2 = pos + 2
			var pos3 = pos + 3

			switch opcode {
			case 99:
				break oploop
			case 1:
				ops[ops[pos3]] = getOpsValue(p1Mode, &ops, pos1) + getOpsValue(p2Mode, &ops, pos2)
				pos += 4
			case 2:
				ops[ops[pos3]] = getOpsValue(p1Mode, &ops, pos1) * getOpsValue(p2Mode, &ops, pos2)
				pos += 4
			case 3:
				ops[ops[pos1]] = getIntCodeInput()
				pos += 2
			case 4:
				putIntCodeOutput(getOpsValue(p1Mode, &ops, pos1))
				pos += 2
			default:
				panic("unkown opcode")
			}

		}

		solution <- sol{noun: n, verb: v, result: ops[0]}
	}
}

func getIntCodeInput() int {
	fmt.Printf("read: %d\n", 0)
	return 1
}

func putIntCodeOutput(output int) {
	fmt.Printf("put: %d\n", output)
}
