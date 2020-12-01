package coprime

import (
	"sort"

	"github.com/lu-dde/adventofcode/internal/coord"
)

// Coprimes is and array of Coprime
type Coprimes struct {
	coprimes       coord.Slice
	limitX, limitY int
	index          int
}

//New Coprimes with limits
func New(limitX, limitY int) *Coprimes {
	return &Coprimes{
		coprimes: coord.Slice{ // starting pairs
			coord.Pair{X: 2, Y: 1},
			coord.Pair{X: 3, Y: 1},
		},

		limitX: limitX,
		limitY: limitY,
		index:  0,
	}
}

func (c *Coprimes) inRange(co coord.Pair) bool {
	return co.X > co.Y && co.X <= c.limitX && co.Y <= c.limitY
}

// GetCoprimes generates coprimes from (0,0) -> (limitN,limitN) in ascending order
// https://en.wikipedia.org/wiki/Coprime_integers#Generating_all_coprime_pairs
func (c *Coprimes) GetCoprimes() coord.Slice {
	for c.next() {
		c.index++
	}

	sort.Stable(&c.coprimes)

	return c.coprimes
}

func (c *Coprimes) next() bool {
	// Branch 1: (2m-n,m)
	// Branch 2: (2m+n,m)
	// Branch 3: (m+2n,n)
	return c.branch1() || c.branch2() || c.branch3()
}

func (c *Coprimes) branch1() bool {
	co := c.coprimes[c.index]
	return c.branchAdd(coord.Pair{X: 2*co.X - co.Y, Y: co.X})
}

func (c *Coprimes) branch2() bool {
	co := c.coprimes[c.index]
	return c.branchAdd(coord.Pair{X: 2*co.X + co.Y, Y: co.X})
}

func (c *Coprimes) branch3() bool {
	co := c.coprimes[c.index]
	return c.branchAdd(coord.Pair{X: 2*co.Y - co.X, Y: co.Y})
}

func (c *Coprimes) branchAdd(co coord.Pair) bool {
	if c.inRange(co) {
		c.coprimes = append(c.coprimes, co)
		return true
	}
	return false
}
