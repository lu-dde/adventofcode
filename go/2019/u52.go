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
	solutions := make(chan int)

	go opscode3(opcodeChan, solutions)

	opcodeChan <- cps
	close(opcodeChan)

	s <- fmt.Sprintf("Solution: %d", <-solutions)

}

func opscode3(opsChan chan []int, solution chan int) {
	var healthcheck = 0

	for ops := range opsChan {
		var pos = 0

	oploop:
		for {
			opcodeCompact := ops[pos]

			opcode := opcodeCompact % 100
			p1Mode := (opcodeCompact / 100) % 10
			p2Mode := (opcodeCompact / 1000) % 10
			//p3Mode := (opcodeCompact / 10000) % 10

			//fmt.Println(ops)
			//fmt.Println(pos, "opcode", opcodeCompact, opcode, p1Mode, p2Mode)

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
				ops[ops[pos1]] = 5
				pos += 2
			case 4:
				healthcheck = getOpsValue(p1Mode, &ops, pos1)
				pos += 2
			case 5:
				if getOpsValue(p1Mode, &ops, pos1) != 0 {
					pos = getOpsValue(p2Mode, &ops, pos2)
				} else {
					pos += 3
				}
			case 6:
				if getOpsValue(p1Mode, &ops, pos1) == 0 {
					pos = getOpsValue(p2Mode, &ops, pos2)
				} else {
					pos += 3
				}
			case 7:
				ops[ops[pos3]] = bool2int(getOpsValue(p1Mode, &ops, pos1) < getOpsValue(p2Mode, &ops, pos2))
				pos += 4
			case 8:
				ops[ops[pos3]] = bool2int(getOpsValue(p1Mode, &ops, pos1) == getOpsValue(p2Mode, &ops, pos2))
				pos += 4
			default:
				panic("unkown opcode")
			}

		}
	}
	solution <- healthcheck
	close(solution)
}

func bool2int(b bool) int {
	if b {
		return 1
	}
	return 0
}
