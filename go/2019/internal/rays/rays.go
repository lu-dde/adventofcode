package rays

import (
	"github.com/lu-dde/adventofcode/internal/coord"
)

type Gen struct {
	Generation int
	Speed      coord.Pair
}

type Map struct {
	m   map[coord.Pair]Gen
	max coord.Pair
}

//Get a Gen struct
func (r *Map) Get(coord coord.Pair) (Gen, bool) {
	g, ok := r.m[coord]
	return g, ok
}

func (r *Map) add(direction coord.Pair) bool {
	_, ok := r.m[direction]
	if ok {
		return false
	}

	generation := 1
	ray := direction
	for ray.Inside(r.max) {
		r.m[ray] = Gen{
			Generation: generation,
			Speed:      direction,
		}
		ray = ray.Add(direction)
		generation++
	}

	return true
}

// New generates a coord map with the speed of the ray
func New(size int) Map {

	r := Map{
		m:   make(map[coord.Pair]Gen, size*size),
		max: coord.NewPair(size, size),
	}

	r.add(coord.NewPair(1, 1))

	for i := 0; i < size; i++ {
		for j := 0; j < i; j++ {
			direction := coord.NewPair(i, j)
			r.add(direction)
			r.add(direction.AsInverted())
		}
	}

	return r
}
