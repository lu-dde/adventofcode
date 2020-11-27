package p2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lu-dde/adventofcode/internal/coord"
	hullpainter "github.com/lu-dde/adventofcode/solutions/u11/HullPainter"
)

//Solve is main proxy for solve, takes a string channel
func Solve(p chan string, s chan string) {

	// we only expect one line.
	line, _ := <-p

	code := strings.Split(line, ",")
	var ops = make([]int64, 1000)

	for index, c := range code {
		i, _ := strconv.ParseInt(c, 10, 64)
		ops[index] = i
	}

	machine := hullpainter.New(ops)

	machine.Hull[machine.Position] = 1

	machine.Run()

	//, what registration identifier does it paint on your hull?

	min, max := getMinMaxCoords(machine.Hull)

	fmt.Println(min, max)
	for row := min.Y; row <= max.Y; row++ {
		for col := min.X; col <= max.X; col++ {
			color, ok := machine.Hull[coord.NewPair(col, row)]

			if ok {
				if color == 1 {
					fmt.Print("#")
				} else {
					fmt.Print(" ")
				}
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	s <- fmt.Sprint("Solution: ", len(machine.Hull))

}

func getMinMaxCoords(m map[coord.Pair]int64) (coord.Pair, coord.Pair) {
	minX, minY := 0, 0
	maxX, maxY := 0, 0

	for co := range m {
		minX = min(minX, co.X)
		minY = min(minY, co.Y)
		maxX = max(maxX, co.X)
		maxY = max(maxY, co.Y)
	}

	return coord.NewPair(minX, minY), coord.NewPair(maxX, maxY)

}

func min(p, q int) int {
	if p < q {
		return p
	}
	return q
}

func max(p, q int) int {
	if p < q {
		return q
	}
	return p
}
