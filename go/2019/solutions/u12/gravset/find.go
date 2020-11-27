package grav

import (
	"fmt"
	"math"
)

//FindCycle find the the step when it's the same state again
func (gs *Set) FindCycle() {
	gs.checkPast()
	for step := 0; step < math.MaxInt64; step++ {
		gs.addVelocity()
		gs.move()
		if gs.checkPast() {
			break
		}
	}
}

func (gs *Set) addVelocity() {
	for i, p := range gs.positions {
		for _, q := range gs.positions {
			gs.velocity[i] += grav(p, q)
		}
	}
}

func (gs *Set) move() {
	for i := range gs.positions {
		gs.positions[i] += gs.velocity[i]
	}
}

func grav(p, o int) int {
	switch {
	case p < o:
		return 1
	case p > o:
		return -1
	default:
		return 0
	}
}

func (gs *Set) checkPast() bool {
	hash := fmt.Sprintf("(%d:%d:%d:%d|%d:%d:%d:%d)",
		gs.positions[0], gs.positions[1], gs.positions[2], gs.positions[3],
		gs.velocity[0], gs.velocity[1], gs.velocity[2], gs.velocity[3],
	)
	if gs.past[hash] {
		fmt.Println("found cycle in axis", gs.axis, "with", len(gs.past), "steps")
		return true
	}
	gs.past[hash] = true
	return false
}
