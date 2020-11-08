package arcade

import (
	"fmt"
	"sync"
	"time"

	"github.com/lu-dde/adventofcode/internal/coord"
	"github.com/lu-dde/adventofcode/internal/intcode"
)

const (
	objVoid   = 0
	objWall   = 1
	objBlock  = 2
	objPaddle = 3
	objBall   = 4

	joystickNet   = 0
	joystickLeft  = -1
	joystickRight = 1
)

var scoreCoord = coord.NewPair(-1, 0)
var zeroCoord = coord.NewPair(0, 0)

//Game arcade struct
type Game struct {
	Screen  *Screen
	Input   chan int64
	Output  chan int64
	machine intcode.V6
	Score   int
	auto    tracker

	started bool
}

// New creates a Arcade Game with an internal intcode.V6 machine
func New(ops []int64) *Game {

	var input = make(chan int64)
	var output = make(chan int64)

	machine := intcode.New(ops, input, output)

	game := Game{
		machine: machine,
		Input:   input,
		Output:  output,
		Screen:  NewScreen(),
		auto:    tracker{},
		started: false,
	}

	return &game
}

//Print game state to terminal
func (game *Game) Print() {
	fmt.Println("Score: ", game.Score)
	//game.Screen.PrintScreen()
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
	ticker := time.NewTicker(300 * time.Microsecond)
	tickerS := time.NewTicker(1000 * time.Millisecond)

	doneJoystick := make(chan bool)
	donePaint := make(chan bool)

	var wg sync.WaitGroup
	go func() {
		wg.Add(1)
		defer wg.Done()
		game.machine.Run()
	}()
	go func() {
		wg.Add(1)
		defer wg.Done()
		for {
			c, oid, more := game.nextPaintInstruction()

			if c == scoreCoord {
				game.Score = oid
			} else {
				game.Screen.Paint(c, oid)
			}

			game.auto.track(c, oid)

			if !more {
				doneJoystick <- true
				donePaint <- true
				break
			}
		}
	}()
	go func() {
		wg.Add(1)
		defer wg.Done()
		for {
			select {
			case <-doneJoystick:
				return
			case <-ticker.C:
				if game.machine.Waiting {
					game.Input <- game.auto.getDirection()
				}
			}
		}
	}()
	go func() {
		wg.Add(1)
		defer wg.Done()
		for {
			select {
			case <-donePaint:
				return
			case <-tickerS.C:
				game.Print()
			}
		}
	}()

	wg.Wait()
}
