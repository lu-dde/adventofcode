package p1

import (
	"fmt"
	"sort"
	"strconv"
)

//Solve is main proxy for solve, takes a string channel
func Solve(p chan string, s chan string) {
	var t = 0

	volts := []int{}

	for line := range p {
		v, _ := strconv.Atoi(line)
		volts = append(volts, v)
	}

	sort.Ints(volts)

	diffs := []int{0, 0, 0, 1}

	current := 0
	for _, voltage := range volts {
		if voltage < current+4 {
			diffs[voltage-current]++
			current = voltage
		} else {
			fmt.Println("break", current, voltage)
			break
		}
	}

	fmt.Println(diffs)

	t = diffs[1] * diffs[3]

	s <- fmt.Sprintf("Solution: %d", t)
}
