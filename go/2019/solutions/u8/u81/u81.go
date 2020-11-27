package u81

import (
	"fmt"
	"strconv"
	"strings"
)

type layer81 struct {
	nr     int
	layer  [][]int
	zeroes int
	ones   int
	twos   int
}

//Solve is main proxy for solve, takes a string channel
func Solve(p chan string, s chan string) {
	// only one line of input
	line, _ := <-p

	reader := strings.NewReader(line)

	wide := 25
	tall := 6
	square := wide * tall

	layerCount := reader.Len() / square

	var layers = make([]layer81, reader.Len()/square)

	var winner = layer81{nr: -1, zeroes: square + 1}

	for layerIndex := 0; layerIndex < layerCount; layerIndex++ {
		// TODO
		var counts = make([]int, 3)
		var partialImage = make([][]int, tall)
		for row := 0; row < tall; row++ {
			partialImage[row] = make([]int, wide)
			for col := 0; col < wide; col++ {
				b, _ := reader.ReadByte() // skip error, we are sure on size
				i, _ := strconv.Atoi(string(b))
				partialImage[row][col] = i
				counts[i]++
			}
		}
		layer := layer81{nr: layerIndex, layer: partialImage, zeroes: counts[0], ones: counts[1], twos: counts[2]}
		layers[layerIndex] = layer
		if layer.zeroes < winner.zeroes {
			winner = layer
		}
	}

	s <- fmt.Sprintf("Solution: %d", winner.ones*winner.twos)
}
