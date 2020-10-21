package coord

//Pair contains an X and an Y coordinate
type Pair struct {
	X, Y int
}

func (cs Pair) less(other Pair) bool {
	return cs.distance() < other.distance()
}

func (cs Pair) distance() int {
	return (cs.X * cs.X) + (cs.Y * cs.Y)
}
