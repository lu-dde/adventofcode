package p2

import "fmt"

const intActive uint8 = 128

//Solve is main proxy for solve, takes a string channel
func Solve(p chan string, s chan string) {
	var t = 0

	cw := cw4d{
		cubes: make(map[Dim]uint8, 128),
	}

	y := 0
	for line := range p {
		for x, a := range line {
			if a == '#' {
				cw.add(Dim{x, y, 0, 0})
			}
		}
		y++
	}

	//	cw.print()
	cw.steps(6)

	t = cw.count()

	s <- fmt.Sprintf("Solution: %d", t)
}

//Dim 4d space
type Dim struct {
	X int
	Y int
	Z int
	W int
}

type cw4d struct {
	cubes map[Dim]uint8
	cycle int
}

func (cw *cw4d) step() {

	nc := make(map[Dim]uint8, intActive)

	for d := range cw.cubes {
		nc[d] += intActive
		for _, n := range cw.neighbours(d) {
			nc[n]++
		}
	}
	cw.cubes = nc

	for d, count := range cw.cubes {
		switch {
		case count == intActive+3:
			nc[d] = intActive
		case count == intActive+2:
			nc[d] = intActive
		case count == 3:
			nc[d] = intActive
		default:
			delete(nc, d)
		}
	}
	cw.cycle++

	//cw.print()
	//fmt.Println()
}

func (cw *cw4d) steps(count int) {
	for i := 0; i < count; i++ {
		cw.step()
	}
}

func (cw *cw4d) add(d Dim) {
	cw.cubes[d] = intActive
}

func (cw *cw4d) count() int {
	count := 0
	for _, active := range cw.cubes {
		if active == intActive {
			count++
		}
	}
	return count
}

func (cw *cw4d) neighbours(c Dim) []Dim {
	st := [3]int{-1, 0, +1}
	neighbours := []Dim{}

	for _, x := range st {
		for _, y := range st {
			for _, z := range st {
				for _, w := range st {
					next := Dim{c.X + x, c.Y + y, c.Z + z, c.W + w}
					if next != c {
						neighbours = append(neighbours, next)
					}
				}
			}
		}
	}

	return neighbours
}

func (cw *cw4d) print() {
	fmt.Println("cycle", cw.cycle)
	fmt.Println("count", cw.count())
	for w := 0; w < 3; w++ {
		fmt.Println("w", w)
		for z := 0; z < 3; z++ {
			fmt.Println("w", w, "z", z)
			for y := -2; y < 5; y++ {
				for x := -2; x < 5; x++ {
					if cw.cubes[Dim{x, y, z, w}] == 0 {
						fmt.Print("\033[1m\033[30m·\033[0m")
					} else if cw.cubes[Dim{x, y, z, w}] > 31 {
						fmt.Printf("\033[1m\033[33m•\033[0m")
					} else {
						fmt.Printf(" %02d", cw.cubes[Dim{x, y, z, w}])
					}
				}
				fmt.Println()
			}
			fmt.Println()
		}
	}
}
