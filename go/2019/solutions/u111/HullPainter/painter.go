package hullpainter

import (
	"fmt"
	"sync"

	"github.com/lu-dde/adventofcode/internal/coord"
	"github.com/lu-dde/adventofcode/solutions/u91"
)

// HullPainter u111
type HullPainter struct {
	machine   u91.Intcode6
	Input     chan int64
	Output    chan int64
	Position  coord.Pair
	Direction coord.Pair
	Hull      map[coord.Pair]pos
}

// New creates a HullPainter with an internal u91.NewIntcode6 machine
func New(ops []int64) *HullPainter {

	var input = make(chan int64, 1)
	var output = make(chan int64)

	//input <- 0

	machine := u91.NewIntcode6(ops, input, output)

	hp := HullPainter{
		machine:   machine,
		Input:     input,
		Output:    output,
		Position:  coordZero,
		Direction: north,
		Hull:      make(map[coord.Pair]pos),
	}

	return &hp
}

//Run the HullPainter
func (hp *HullPainter) Run() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		hp.machine.Run()
	}()
	go func() {
		defer wg.Done()
		hp.paint()
	}()
	wg.Wait()
}

func (hp *HullPainter) paint() {
	for {
		hp.Input <- hp.Hull[hp.Position].color

		paint, ok := <-hp.Output
		if !ok {
			fmt.Println("end of paint")
			break
		}
		rotateCmd := <-hp.Output

		hp.Hull[hp.Position] = pos{
			color:     paint,
			direction: hp.Direction,
		}

		hp.Direction = rotate(rotateCmd, hp.Direction)
		hp.Position = hp.Position.Add(hp.Direction)

	}
}
