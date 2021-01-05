package p1

import (
	"fmt"
)

//Solve is main proxy for solve, takes a string channel
func Solve(p chan string, s chan string) {
	var t = 0

	cw := cw3d{
		cubes: make(map[Dim]uint8, 256),
	}

	y := 0
	for line := range p {
		for x, a := range line {
			if a == '#' {
				cw.add(Dim{x, y, 0})
			}
		}
		y++
	}

	//cw.print()
	cw.steps(6)

	t = cw.count()

	s <- fmt.Sprintf("Solution: %d", t)
}

//Dim 3d space
type Dim struct {
	X int
	Y int
	Z int
}

type cw3d struct {
	cubes map[Dim]uint8
	cycle int
}

func (cw *cw3d) step() {

	nc := make(map[Dim]uint8, 256)

	for d := range cw.cubes {
		nc[d] += 32
		for _, n := range cw.neighbours(d) {
			nc[n]++
		}
	}
	cw.cubes = nc

	for d, count := range cw.cubes {
		switch {
		case count == 34:
			nc[d] = 32
		case count == 35:
			nc[d] = 32
		case count == 3:
			nc[d] = 32
		default:
			delete(nc, d)
		}
	}
	cw.cycle++

	//cw.print()
	//fmt.Println()
}

func (cw *cw3d) steps(count int) {
	for i := 0; i < count; i++ {
		cw.step()
	}
}

func (cw *cw3d) add(d Dim) {
	cw.cubes[d] = 32
}

func (cw *cw3d) count() int {
	count := 0
	for _, active := range cw.cubes {
		if active == 32 {
			count++
		}
	}
	return count
}

func (cw *cw3d) neighbours(c Dim) []Dim {
	st := [3]int{-1, 0, +1}
	neighbours := []Dim{}

	for _, x := range st {
		for _, y := range st {
			for _, z := range st {
				next := Dim{c.X + x, c.Y + y, c.Z + z}
				if next != c {
					neighbours = append(neighbours, next)
				}
			}
		}
	}

	return neighbours
}

func (cw *cw3d) print() {
	fmt.Println("cycle", cw.cycle)
	fmt.Println("count", cw.count())
	for z := 0; z < 3; z++ {
		fmt.Println("z", z)
		for y := -2; y < 5; y++ {
			for x := -2; x < 5; x++ {
				if cw.cubes[Dim{x, y, z}] == 0 {
					fmt.Print(" \033[1m\033[30m·\033[0m ")
				} else if cw.cubes[Dim{x, y, z}] > 31 {
					fmt.Printf(" \033[1m\033[33m%02d\033[0m", cw.cubes[Dim{x, y, z}])
				} else {
					fmt.Printf(" %02d", cw.cubes[Dim{x, y, z}])
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}

}
