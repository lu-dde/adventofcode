package hullpainter

import "github.com/lu-dde/adventofcode/internal/coord"

var (
	north     = coord.Pair{X: 0, Y: -1}
	south     = coord.Pair{X: 0, Y: 1}
	west      = coord.Pair{X: -1, Y: 0}
	east      = coord.Pair{X: 1, Y: 0}
	coordZero = coord.Pair{X: 0, Y: 0}
)

func rotateLeft(p coord.Pair) coord.Pair {
	switch p {
	case north:
		return west
	case west:
		return south
	case south:
		return east
	case east:
		return north
	}
	return coordZero

}

func rotateRight(p coord.Pair) coord.Pair {
	switch p {
	case north:
		return east
	case east:
		return south
	case south:
		return west
	case west:
		return north
	}
	return coordZero
}

func rotate(rotate int64, p coord.Pair) coord.Pair {
	if rotate == 0 {
		return rotateLeft(p)
	} else if rotate == 1 {
		return rotateRight(p)
	}
	return coordZero
}
