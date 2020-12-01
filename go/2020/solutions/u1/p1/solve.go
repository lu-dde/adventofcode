package p1

import (
	"fmt"
	"strconv"
)

//Solve1 is main proxy for solve, takes a string channel
func Solve1(p chan string, s chan string) {
	var t = 0

	var want map[int]int = make(map[int]int, 200)

	for line := range p {
		i, _ := strconv.Atoi(line)

		if want[i] > 0 {
			t = want[i] * i
			fmt.Println(want[i], "*", i, "=", want[i]*i)
		} else {
			want[2020-i] = i

		}
	}

	s <- fmt.Sprintf("Solution: %d", t)
}
