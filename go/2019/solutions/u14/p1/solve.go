package p1

import (
	"fmt"

	"github.com/lu-dde/adventofcode/solutions/u14/alch"
)

//Solve func
func Solve(p chan string, s chan string) {

	synth := alch.New("FUEL")

	for line := range p {
		synth.AddLine(line)
	}

	cost := synth.FuelCost()

	s <- fmt.Sprintf("%d ORE for 1 FUEL", cost)
}
