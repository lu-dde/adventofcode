package u82

import (
	"fmt"
	"strings"
)

//Solve is main proxy for solve, takes a string channel
func Solve(p chan string, s chan string) {
	// only one line of input
	line, _ := <-p

	reader := strings.NewReader(line)

	wide := 25
	tall := 6
	square := wide * tall

	layerCount := reader.Len() / square

	merge := func(old, new byte) byte {
		if old == '2' {
			return new
		}
		return old
	}

	toPrint := func(b byte) string {
		if b == '1' {
			return "•"
		}
		return " "
	}

	var image = make([]byte, square)
	for index := 0; index < len(image); index++ {
		image[index] = '2'
	}

	for layerIndex := 0; layerIndex < layerCount; layerIndex++ {
		for rc := 0; rc < square; rc++ {
			b, _ := reader.ReadByte()
			image[rc] = merge(image[rc], b)
		}
	}

	for row := 0; row < tall; row++ {
		for col := 0; col < wide; col++ {
			offset := row*wide + col
			fmt.Print(toPrint(image[offset]))
		}
		fmt.Println()
	}

	s <- fmt.Sprintf("Solution: %d", square)
}
