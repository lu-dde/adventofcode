package hullpainter

import (
	"sync"

	"github.com/lu-dde/adventofcode/internal/coord"
	"github.com/lu-dde/adventofcode/solutions/u9p/u91"
)

// HullPainter u111
type HullPainter struct {
	machine   u91.Intcode6
	Input     chan int64
	Output    chan int64
	Position  coord.Pair
	Direction coord.Pair
	Hull      map[coord.Pair]int64
}

// New creates a HullPainter with an internal u91.NewIntcode6 machine
func New(ops []int64) *HullPainter {

	var input = make(chan int64, 1)
	var output = make(chan int64)

	machine := u91.NewIntcode6(ops, input, output)

	hp := HullPainter{
		machine:   machine,
		Input:     input,
		Output:    output,
		Position:  coordZero,
		Direction: north,
		Hull:      make(map[coord.Pair]int64),
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

func (hp *HullPainter) getNextCommand() (paint, rotateCmd int64, ok bool) {
	paint, ok = <-hp.Output
	if !ok {
		return 0, 0, false
	}
	rotateCmd = <-hp.Output

	return paint, rotateCmd, true
}

func (hp *HullPainter) paintStep() bool {
	hp.Input <- hp.Hull[hp.Position]

	paint, rotateCmd, ok := hp.getNextCommand()
	if !ok {
		return false
	}

	hp.Hull[hp.Position] = paint

	hp.Direction = rotate(rotateCmd, hp.Direction)
	hp.Position = hp.Position.Add(hp.Direction)

	return true
}
func (hp *HullPainter) paint() {
	for hp.paintStep() {
	}
}
