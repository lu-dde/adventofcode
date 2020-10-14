package u51

import (
	"fmt"
	"strconv"
	"strings"
)

//Solve is main proxy for solve, takes a string channel
func Solve(p chan string, s chan string) {

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

type sol struct {
	noun   int
	verb   int
	result int
}

//GetOpsValue get operation int value
func GetOpsValue(mode int, ops *[]int, pos int) int {
	if mode == 0 {
		source1 := (*ops)[pos]
		return (*ops)[source1]
	} else if mode == 1 {
		return (*ops)[pos]
	}
	panic("unkown mode")
}

func opscode2(opsChan chan []int, solution chan sol) {

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
				ops[ops[pos3]] = GetOpsValue(p1Mode, &ops, pos1) + GetOpsValue(p2Mode, &ops, pos2)
				pos += 4
			case 2:
				ops[ops[pos3]] = GetOpsValue(p1Mode, &ops, pos1) * GetOpsValue(p2Mode, &ops, pos2)
				pos += 4
			case 3:
				ops[ops[pos1]] = 1
				pos += 2
			case 4:
				fmt.Printf("put: %d\n", GetOpsValue(p1Mode, &ops, pos1))
				pos += 2
			default:
				panic("unkown opcode")
			}

		}

		solution <- sol{noun: n, verb: v, result: ops[0]}
	}
}
