package p2

import (
	"fmt"
	"strconv"
)

//Solve is main proxy for solve, takes a string channel
func Solve(p chan string, s chan string) {
	var t uint64 = 0

	var weakness uint64 = 1492208709

	length := 660
	var sum uint64 = 0
	sums := make([]uint64, length)
	for i := range sums {
		v, _ := strconv.ParseUint(<-p, 10, 64)
		sum += v
		sums[i] = sum
		if v > weakness {
			break
		}
	}

	var I int = 0
	var J int = 0

outer:
	for i := length - 1; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if sums[i]-sums[j] == weakness {
				I, J = i, j
				break outer
			} else if sums[i]-sums[j] < weakness {
				continue outer
			}
		}

	}

	var min uint64 = sums[I+1] - sums[I]
	var max uint64 = min
	for i := J; i < I; i++ {
		current := sums[i+1] - sums[i]
		if current < min {
			min = current
		}
		if current > max {
			max = current
		}
	}
	t = min + max

	s <- fmt.Sprintf("Solution: %d", t)
}
