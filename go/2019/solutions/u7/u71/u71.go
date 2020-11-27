package u71

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lu-dde/adventofcode/solutions/u5/u51"
)

//Solve is main proxy for solve, takes a string channel
func Solve(p chan string, s chan string) {

	// we only expect one line.
	line, _ := <-p

	code := strings.Split(line, ",")
	var ops = []int{}

	for _, c := range code {
		i, _ := strconv.Atoi(c)
		ops = append(ops, i)
	}

	perms := MakeAmpPerms([]int{0, 1, 2, 3, 4})

	var winnerCFG []int
	var winnerB = 0
	for _, cfg := range perms {
		var buffer = 0
		for _, c := range cfg {
			buffer = opscode4(ops, []int{c, buffer})
		}
		if winnerB < buffer {
			winnerB = buffer
			winnerCFG = cfg
			fmt.Println(cfg, buffer)
		}
	}

	s <- fmt.Sprint("Solution: ", winnerB, winnerCFG)

}

//MakeAmpPerms amplification permutaions
func MakeAmpPerms(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func opscode4(ops []int, phaseSettings []int) int {
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
			ops[ops[pos3]] = u51.GetOpsValue(p1Mode, &ops, pos1) + u51.GetOpsValue(p2Mode, &ops, pos2)
			pos += 4
		case 2:
			ops[ops[pos3]] = u51.GetOpsValue(p1Mode, &ops, pos1) * u51.GetOpsValue(p2Mode, &ops, pos2)
			pos += 4
		case 3:
			ops[ops[pos1]] = phaseSettings[0]
			phaseSettings = phaseSettings[1:]
			pos += 2
		case 4:
			healthcheck = u51.GetOpsValue(p1Mode, &ops, pos1)
			pos += 2
		case 5:
			if u51.GetOpsValue(p1Mode, &ops, pos1) != 0 {
				pos = u51.GetOpsValue(p2Mode, &ops, pos2)
			} else {
				pos += 3
			}
		case 6:
			if u51.GetOpsValue(p1Mode, &ops, pos1) == 0 {
				pos = u51.GetOpsValue(p2Mode, &ops, pos2)
			} else {
				pos += 3
			}
		case 7:
			ops[ops[pos3]] = bool2int(u51.GetOpsValue(p1Mode, &ops, pos1) < u51.GetOpsValue(p2Mode, &ops, pos2))
			pos += 4
		case 8:
			ops[ops[pos3]] = bool2int(u51.GetOpsValue(p1Mode, &ops, pos1) == u51.GetOpsValue(p2Mode, &ops, pos2))
			pos += 4
		default:
			panic("unkown opcode")
		}

	}

	return healthcheck
}

func bool2int(b bool) int {
	if b {
		return 1
	}
	return 0
}
