package coord

//Pair contains an X and an Y coordinate
type Pair struct {
	X, Y int
}

//Less returns true if other is farther away from (0,0) then cs
func (cs Pair) Less(other Pair) bool {
	return cs.Score() < other.Score()
}

//Score is calculated by X^2 + Y^2
func (cs Pair) Score() int {
	return (cs.X * cs.X) + (cs.Y * cs.Y)
}
