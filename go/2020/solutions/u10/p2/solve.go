package p2

import (
	"fmt"
	"sort"
	"strconv"
)

//Solve is main proxy for solve, takes a string channel
func Solve(p chan string, s chan string) {
	var t int64 = 0

	volts := []int{0}

	for line := range p {
		v, _ := strconv.Atoi(line)
		volts = append(volts, v)
	}

	sort.Ints(volts)

	current := 0
	for _, voltage := range volts {
		if voltage < current+4 {
			current = voltage
		}
	}
	fmt.Println("break", current)

	vlen := len(volts)

	//        	 steps nr-of-paths
	tree := make([]int64, vlen)

	tree[0] = 1

	for i := 0; i < vlen; i++ {
		volt := volts[i]
		paths := tree[i]
		for j := i + 1; j < i+4 && j < vlen; j++ {
			volt2 := volts[j]
			if volt2 < volt+4 {
				tree[j] += paths
			}
		}
	}

	t = tree[vlen-1]

	s <- fmt.Sprintf("Solution: %d", t)
}
