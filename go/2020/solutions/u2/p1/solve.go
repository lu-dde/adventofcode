package p1

import (
	"fmt"

	"github.com/lu-dde/adventofcode/solutions/u2/policy"
)

//Solve1 is main proxy for solve, takes a string channel
func Solve1(p chan string, s chan string) {
	var t = 0

	for line := range p {
		pass := policy.ParsePass(line)
		if pass.Count >= pass.Min && pass.Count <= pass.Max {
			//fmt.Println(pass)
			t++
		}
	}

	s <- fmt.Sprintf("Solution: %d", t)
}
