package p2

import (
	"fmt"
	"strings"
)

//Solve is main proxy for solve, takes a string channel
func Solve(p chan string, s chan string) {
	var t = 0

	nums := strings.Split(<-p, ",")

	t = len(nums)

	s <- fmt.Sprintf("Solution: %d", t)
}
