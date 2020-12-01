package coord

//NewPair creates a new pair X,Y
func NewPair(x, y int) Pair {
	return Pair{X: x, Y: y}
}

//Pair contains an X and an Y coordinate
type Pair struct {
	X, Y int
}

//Less returns true if other is farther away from (0,0) then cs
func (p Pair) Less(other Pair) bool {
	return p.Score() < other.Score()
}

//Inside returns true if other strictly outside of p
func (p Pair) Inside(other Pair) bool {
	return p.X < other.X && p.Y < other.Y
}

//Score is calculated by X^2 + Y^2
func (p Pair) Score() int {
	return (p.X * p.X) + (p.Y * p.Y)
}

//AsInverted returns a new Pair with X and Y swapped
func (p Pair) AsInverted() Pair {
	return Pair{X: p.Y, Y: p.X}
}

//Add two Pairs togheter and return a new Pair
func (p Pair) Add(other Pair) Pair {
	return Pair{X: p.X + other.X, Y: p.Y + other.Y}
}

//Distance takes the distance between p and other
func (p Pair) Distance(other Pair) (distance Pair, direction Pair) {
	x := p.X - other.X
	y := p.Y - other.Y

	polarityX := 1
	if x < 0 {
		polarityX = -1
	}

	polarityY := 1
	if y < 0 {
		polarityY = -1
	}

	return NewPair(x*polarityX, y*polarityY), NewPair(polarityX, polarityY)
}

//ChangeDirection of p to be the same as direction
func (p Pair) ChangeDirection(direction Pair) Pair {
	return NewPair(p.X*direction.X, p.Y*direction.Y)
}
