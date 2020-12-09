package p2

import (
	"fmt"
	"strconv"
)

//var weaknessPos int = 14

var weaknessPos int = 659

//Solve is main proxy for solve, takes a string channel
func Solve(p chan string, s chan string) {
	var t uint64 = 0

	length := weaknessPos + 1
	loop := make([]uint64, length)
	for i := range loop {
		v, _ := strconv.ParseUint(<-p, 10, 64)
		loop[i] = v
	}

	//	fmt.Println(loop)
	var weakness uint64 = loop[weaknessPos]

outer:
	for i := 0; i < length; i++ {
		sum := loop[i]
		min := sum
		max := sum

		for j := i + 1; j < length; j++ {
			cur := loop[j]
			sum += cur
			if min > cur {
				min = cur
			}
			if max < cur {
				max = cur
			}
			if sum == weakness {
				t = min + max
				break outer
			}
		}

	}

	s <- fmt.Sprintf("Solution: %d", t)
}
