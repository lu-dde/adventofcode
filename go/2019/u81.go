package main

import (
	"fmt"
	"strconv"
	"strings"
)

//U81 is main proxy for solve, takes a string channel
func U81(p chan string, s chan string) {
	// only one line of input
	line, _ := <-p

	reader := strings.NewReader(line)

	wide := 25
	tall := 3

	var layers = make([][]int, tall)

	for row := 0; row < tall; row++ {
		layers[row] = make([]int, wide)
		for col := 0; col < wide; col++ {
			b, _ := reader.ReadByte() // skip error, we are sure on size
			i, _ := strconv.Atoi(string(b))
			layers[row][col] = i
		}
	}

	s <- fmt.Sprintf("Solution: %d", layers)
}
