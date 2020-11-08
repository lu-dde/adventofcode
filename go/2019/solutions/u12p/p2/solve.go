package p2

import (
	"fmt"
	"sync"

	"github.com/lu-dde/adventofcode/internal/gcd"
	grav "github.com/lu-dde/adventofcode/solutions/u12p/gravset"
)

//Solve is main proxy for solve, takes a string channel
func Solve(p chan string, s chan string) {

	o := orbits{
		axis: []*grav.Set{
			grav.New(0),
			grav.New(1),
			grav.New(2),
		},
	}

	for line := range p {
		fmt.Println(line)
		o.parseStr(line)
	}

	o.run()

	steps := o.steps()

	s <- fmt.Sprint("The number of steps taken: ", steps)

}

type orbits struct {
	stepIndex int
	axis      []*grav.Set
}

func (o *orbits) parseStr(line string) {
	for _, s := range o.axis {
		s.ParseStr(line)
	}
}

func (o *orbits) run() {
	var wg sync.WaitGroup
	for _, s := range o.axis {
		wg.Add(1)
		go func(s *grav.Set) {
			defer wg.Done()
			s.FindCycle()
		}(s)
	}
	wg.Wait()
}

func (o *orbits) steps() int {
	m := []int{0, 0, 0}
	for i, s := range o.axis {
		m[i] = s.Steps()
	}

	return gcd.LCM(m[0], m[1], m[2])
}
