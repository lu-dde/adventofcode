package p1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lu-dde/adventofcode/solutions/u13p/arcade"
)

//Solve func
func Solve(p chan string, s chan string) {

	// we only expect one line.
	line, _ := <-p

	code := strings.Split(line, ",")
	var ops = make([]int64, 4096)

	for index, c := range code {
		i, _ := strconv.ParseInt(c, 10, 64)
		ops[index] = i
	}

	machine := arcade.New(ops)

	machine.Run()

	machine.Print()

	count := machine.Screen.CountBlocks()

	s <- fmt.Sprint("answer: ", count)

}
