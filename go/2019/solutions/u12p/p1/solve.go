package p1

import (
	"fmt"

	"github.com/lu-dde/adventofcode/solutions/u12p/moon"
)

//Solve is main proxy for solve, takes a string channel
func Solve(p chan string, s chan string) {

	o := orbits{
		moons: []*moon.Moon{},
	}

	// we only expect one line.
	for line := range p {
		fmt.Println(line)
		moon := moon.ParseStr(line)
		o.moons = append(o.moons, moon)
	}

	o.printState()
	for i := 0; i < 1000; i++ {
		o.step()
	}
	o.printState()

	fmt.Println("Sum of total energy: ", o.energy())

	s <- fmt.Sprint("Sum of total energy: ", o.energy())

}

type orbits struct {
	moons     []*moon.Moon
	stepIndex int
}

func (o *orbits) gravitate() {
	for i, m := range o.moons {
		//m.Reset()
		for j, o := range o.moons {
			if i != j {
				m.Gravitate(o)
			}
		}
	}
}

func (o *orbits) move() {
	for _, m := range o.moons {
		m.Move()
	}
}

func (o *orbits) step() {
	o.stepIndex++
	o.gravitate()
	o.move()
}

func (o *orbits) energy() float64 {
	var sum float64
	for _, m := range o.moons {
		sum += m.Energy()
	}

	return sum
}

func (o *orbits) printState() {
	fmt.Printf("After %d steps:", o.stepIndex)
	fmt.Println()
	for _, m := range o.moons {
		fmt.Println(*m)
	}
}
