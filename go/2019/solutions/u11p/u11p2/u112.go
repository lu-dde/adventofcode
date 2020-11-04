package u11p2

import (
	"fmt"
	"strconv"
	"strings"

	hullpainter "github.com/lu-dde/adventofcode/solutions/u11p/HullPainter"
)

//Solve is main proxy for solve, takes a string channel
func Solve(p chan string, s chan string) {

	// we only expect one line.
	line, _ := <-p

	code := strings.Split(line, ",")
	var ops = make([]int64, 4096) // TODO

	for index, c := range code {
		i, _ := strconv.ParseInt(c, 10, 64)
		ops[index] = i
	}

	machine := hullpainter.New(ops)

	machine.Hull[machine.Position] = 1

	machine.Run()

	s <- fmt.Sprint("Solution: ", len(machine.Hull))

}
