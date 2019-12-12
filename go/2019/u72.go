package main

import (
	"fmt"
	"strconv"
	"strings"
)

//U72 is main proxy for solve, takes a string channel
func U72(p chan string, s chan string) {

	// we only expect one line.
	line, _ := <-p

	code := strings.Split(line, ",")
	var ops = []int{}

	for _, c := range code {
		i, _ := strconv.Atoi(c)
		ops = append(ops, i)
	}

	perms := makeAmpPerms([]int{5, 6, 7, 8, 9})

	var winnerCFG []int
	var winnerB = 0
	for _, cfg := range perms {

		var feed = map[int](chan int){}
		var cps = map[int]([]int){}
		for index := 0; index < 5; index++ {
			feed[index] = make(chan int)
			cps[index] = make([]int, len(ops))
			copy(cps[index], ops)
		}

		var feedback = make(chan int)

		go opscode5(feed[0], feed[1], cps[0], 0)
		go opscode5(feed[1], feed[2], cps[1], 1)
		go opscode5(feed[2], feed[3], cps[2], 2)
		go opscode5(feed[3], feed[4], cps[3], 3)
		go opscode5(feed[4], feedback, cps[4], 4)

		for index := 0; index < 5; index++ {
			feed[index] <- cfg[index]
		}

		feed[0] <- 0 // init first buffer

		go func(cfg []int) {
			for {
				f, ok := <-feedback
				if ok {
					if winnerB < f {
						winnerB = f
						winnerCFG = cfg
					}
					feed[0] <- f
				}
			}
		}(cfg)
	}

	s <- fmt.Sprint("Solution: ", winnerB, winnerCFG)

}

func opscode5(input chan int, output chan int, ops []int, id int) {
	var healthcheck = 0
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
			in, ok := <-input
			if ok {
				ops[ops[pos1]] = in
				//fmt.Println(id, " got input ", ops[ops[pos1]])
			} else {
				fmt.Println(id, " input chan FAILED")
				break
			}
			pos += 2
		case 4:
			healthcheck = getOpsValue(p1Mode, &ops, pos1)
			//fmt.Println(id, " sending ", healthcheck)
			output <- healthcheck
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
			fmt.Println(id, pos, "opcode", opcodeCompact, opcode, p1Mode, p2Mode)
			panic("unkown opcode")
		}

	}
	//fmt.Println(id, " closing opcode ", healthcheck)
	close(output)
}
