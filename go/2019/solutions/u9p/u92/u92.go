package u92

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lu-dde/adventofcode/solutions/u9p/u91"
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

	var input = make(chan int64, 1)
	var output = make(chan int64)
	var solution = make(chan int64)

	input <- 2

	machine := u91.NewIntcode6(ops, input, output)

	go func() {
		machine.Run()
	}()

	go func() {
		var out int64
		for o := range output {
			out = o
		}
		solution <- out
	}()

	s <- fmt.Sprint("Solution: ", <-solution)

}
