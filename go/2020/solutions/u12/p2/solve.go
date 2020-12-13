package p2

import (
	"fmt"
	"math"
	"strconv"
)

//Solve is main proxy for solve, takes a string channel
func Solve(p chan string, s chan string) {
	var t = 0

	x, y := 0, 0
	wx, wy := 10, 1

	for line := range p {
		d := line[0]
		amount, _ := strconv.Atoi(line[1:])

		switch d {
		case 'F':
			x, y = x+(wx*amount), y+(wy*amount)
		case 'L':
			wx, wy = right(wx, wy, 360-amount)
		case 'R':
			wx, wy = right(wx, wy, amount)
		case 'N':
			wy += amount
		case 'S':
			wy -= amount
		case 'E':
			wx += amount
		case 'W':
			wx -= amount
		}

	}

	t = int(math.Abs(float64(x)) + math.Abs(float64(y)))

	s <- fmt.Sprintf("Solution: %d", t)
}

func right(x, y, deg int) (int, int) {
	t := int(deg / 90)

	for i := 0; i < t; i++ {
		x, y = y, x*-1
	}

	return x, y
}
