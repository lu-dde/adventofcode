package u1

import (
	"fmt"
	"strconv"
)

//Solve2 is main proxy for solve, takes a string channel
func Solve2(p chan string, s chan string) {
	var t = 0

	for {
		line, ok := <-p
		if ok {
			i, _ := strconv.Atoi(line)
			t += getFuel(i)
		} else {
			break
		}
	}

	s <- fmt.Sprintf("Solution: %d", t)
}

func getFuel(i int) int {
	if i < 9 {
		return 0
	}

	r := i/3 - 2
	return r + getFuel(r)

}
