package main

import (
	"fmt"
	"strconv"
	"strings"
)

//U92 is main proxy for solve, takes a string channel
func U92(p chan string, s chan string) {

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

	machine := intcode6{
		ops:    ops,
		input:  input,
		output: output,
	}

	go func() {
		machine.run()
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
