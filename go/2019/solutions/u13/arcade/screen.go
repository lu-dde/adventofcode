package arcade

import (
	"fmt"

	"github.com/lu-dde/adventofcode/internal/coord"
)

/*
0 is an empty tile. No game object appears in this tile.
1 is a wall tile. Walls are indestructible barriers.
2 is a block tile. Blocks can be broken by the ball.
3 is a horizontal paddle tile. The paddle is indestructible.
4 is a ball tile. The ball moves diagonally and bounces off objects.
*/
var printMap = map[int]string{
	objVoid:   " ",
	objWall:   "█",
	objBlock:  "░",
	objPaddle: "—",
	objBall:   "•",
}

//Screen representation
type Screen struct {
	objects map[coord.Pair]int
	cols    int
	rows    int
}

//NewScreen init a new Screen
func NewScreen() *Screen {
	return &Screen{
		objects: map[coord.Pair]int{},
	}
}

func max(p, q int) int {
	if p < q {
		return q
	}
	return p
}

//Paint an object into memory
func (sc *Screen) Paint(c coord.Pair, objectID int) {
	sc.objects[c] = objectID
	sc.cols = max(sc.cols, c.X)
	sc.rows = max(sc.rows, c.Y)
}

//PrintScreen in terminal
func (sc *Screen) PrintScreen() {
	for y := 0; y <= sc.rows; y++ {
		for x := 0; x <= sc.cols; x++ {
			fmt.Print(asPrint(sc.objects[coord.NewPair(x, y)]))
		}
		fmt.Println()
	}
}

func asPrint(o int) string {
	return printMap[o]
}

//CountBlocks the number of block objects on screen
func (sc *Screen) CountBlocks() int {
	count := 0
	for _, oid := range sc.objects {
		if oid == objBlock {
			count++
		}
	}
	return count
}
