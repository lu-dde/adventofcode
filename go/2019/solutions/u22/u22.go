package u22

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

	opcodeChan := make(chan []int, 100*100)
	solutions := make(chan sol, 100*100)

	go opscode(opcodeChan, solutions)

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			var ops = make([]int, len(cps))
			copy(ops, cps)

			ops[1] = noun
			ops[2] = verb

			opcodeChan <- ops
		}
	}

	for solution := range solutions {
		if solution.result == 19690720 {
			s <- fmt.Sprintf("Solution: %d, %v", 100*solution.noun+solution.verb, solution)
			break
		}
	}

}

type sol struct {
	noun   int
	verb   int
	result int
}

func opscode(opsChan chan []int, solution chan sol) {

	for ops := range opsChan {
		n := ops[1]
		v := ops[2]

		step := 4
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

		solution <- sol{noun: n, verb: v, result: ops[0]}
	}

}
