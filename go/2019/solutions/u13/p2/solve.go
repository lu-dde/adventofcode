package p2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lu-dde/adventofcode/solutions/u13/arcade"
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

	ops[0] = 2

	machine := arcade.New(ops)

	machine.Run()

	s <- fmt.Sprint("answer: ", machine.Score)

}
