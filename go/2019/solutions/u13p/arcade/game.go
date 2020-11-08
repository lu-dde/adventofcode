package arcade

import (
	"fmt"
	"sync"

	"github.com/lu-dde/adventofcode/internal/coord"
	"github.com/lu-dde/adventofcode/internal/intcode"
)

var scoreCoord = coord.NewPair(-1, 0)

//Game arcade struct
type Game struct {
	Screen  *Screen
	Input   chan int64
	Output  chan int64
	machine intcode.V6
	score   int
}

// New creates a Arcade Game with an internal intcode.V6 machine
func New(ops []int64) *Game {

	var input = make(chan int64, 1)
	var output = make(chan int64)

	machine := intcode.New(ops, input, output)

	game := Game{
		machine: machine,
		Input:   input,
		Output:  output,
		Screen:  NewScreen(),
	}

	return &game
}

//Print game state to terminal
func (game *Game) Print() {
	fmt.Println("Score: ", game.score)
	game.Screen.PrintScreen()
}

func (game *Game) nextPaintInstruction() (c coord.Pair, objectID int, more bool) {
	x := <-game.Output
	y := <-game.Output
	c = coord.NewPair(int(x), int(y))
	oid, ok := <-game.Output
	return c, int(oid), ok
}

//Run the internal machine and paint
func (game *Game) Run() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		game.machine.Run()
	}()
	go func() {
		defer wg.Done()
		for {
			c, oid, more := game.nextPaintInstruction()

			if c == scoreCoord {
				game.score = oid
				game.Print()
			} else {
				game.Screen.Paint(c, oid)
			}

			if !more {
				break
			}
		}
	}()
	wg.Wait()
}
