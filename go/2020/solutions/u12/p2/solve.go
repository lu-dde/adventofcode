package p2

import (
	"fmt"
	"math"
	"strconv"
)

//Solve is main proxy for solve, takes a string channel
func Solve(p chan string, s chan string) {
	var t = 0

	x := 0.0
	y := 0.0

	var dir byte = 'E'

	for line := range p {
		d := line[0]
		i, _ := strconv.ParseFloat(line[1:], 64)

		switch d {
		case 'F':
			d = dir
		case 'L':
			dir = right(dir, 360-i)
			d = dir
			i = 0
		case 'R':
			dir = right(dir, i)
			d = dir
			i = 0
		}

		x, y = move(d, i, x, y)

	}

	t = int(math.Abs(x) + math.Abs(y))

	s <- fmt.Sprintf("Solution: %d", t)
}

func right(d byte, deg float64) byte {
	t := int(deg / 90)

	for i := 0; i < t; i++ {
		switch d {
		case 'N':
			d = 'E'
		case 'E':
			d = 'S'
		case 'S':
			d = 'W'
		case 'W':
			d = 'N'
		}
	}

	return d
}

func move(d byte, i, x, y float64) (float64, float64) {
	switch d {
	case 'N':
		x += i
	case 'S':
		x -= i
	case 'W':
		y -= i
	case 'E':
		y += i
	}

	return x, y
}
