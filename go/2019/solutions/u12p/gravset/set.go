package grav

//Set contains positions and velocity in an axis
type Set struct {
	axis      int
	positions []int
	velocity  []int
	past      map[string]bool
}

//New returns a new Set
func New(axis int) *Set {
	return &Set{
		axis:      axis,
		positions: []int{},
		velocity:  []int{},
		past:      map[string]bool{},
	}
}

//Steps until cycle
func (gs *Set) Steps() int {
	return len(gs.past)
}
