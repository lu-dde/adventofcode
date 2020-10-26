package u101

import (
	"fmt"

	"github.com/lu-dde/adventofcode/internal/coord"
	"github.com/lu-dde/adventofcode/internal/rays"
)

const debugPrint = false

//Solve U101
func Solve(p chan string, s chan string) {

	size := 40

	stars := starmap{
		chart: make(map[coord.Pair]int, size*size),
	}

	/* starmap '#' is an astroid
	#.#...#.#.
	.###....#.
	.#....#...
	##.#.#.#.#
	....#.#.#.
	.##..###.#
	..#...##..
	..##....##
	......#...
	.####.###.
	*/
	y := 0
	for line := range p {
		for x, t := range line {
			if t == '#' {
				stars.addAstroid(asCoord(x, y))
			}
		}
		y++
	}

	stars.setRayMap()

	if debugPrint {
		for row := 0; row < stars.maxY; row++ {
			for column := 0; column < stars.maxX; column++ {
				ray := asCoord(row, column)
				distance, ok := stars.rays.Get(ray)
				if ok && ray == distance.Speed {
					fmt.Print("•")
				} else if ok {
					fmt.Print("•")
				} else {
					fmt.Print("Ω")
				}
			}
			fmt.Println()
		}
		fmt.Println()

		stars.printBoard()
	}

	stars.look()
	//stars.printBoard()

	var maxCoord coord.Pair
	maxValue := 0
	for coord, value := range stars.chart {
		if value > maxValue {
			maxValue = value
			maxCoord = coord
		}
	}

	s <- fmt.Sprintf("Best is %d,%d with %d other asteroids detected.", maxCoord.X, maxCoord.Y, maxValue)
}

func asCoord(x, y int) coord.Pair {
	return coord.NewPair(x, y)
}

type starmap struct {
	chart map[coord.Pair]int
	maxX  int
	maxY  int
	rays  rays.Map
}

func (s *starmap) addAstroid(coord coord.Pair) {
	s.chart[coord] = 0
	s.extend(coord)
}

func (s *starmap) extend(coord coord.Pair) {
	if s.maxX < coord.X {
		s.maxX = coord.X
	}
	if s.maxY < coord.Y {
		s.maxY = coord.Y
	}
}

func (s *starmap) setRayMap() {
	small, large := s.maxX, s.maxY
	if small > large {
		small, large = large, small
	}
	s.rays = rays.New(large)
}

func (s *starmap) look() {
	for a := range s.chart {
		s.lookNorth(a)
	}
}

func (s *starmap) clearPath(goal, other coord.Pair) bool {
	distance, direction := goal.Distance(other)
	ray, _ := s.rays.Get(distance)
	speed := ray.Speed.ChangeDirection(direction)

	for {
		other = other.Add(speed)
		if goal == other {
			break
		}
		if _, ok := s.chart[other]; ok {
			return false
		}
	}

	return true
}

func (s *starmap) lookNorth(coord coord.Pair) {
	for a := range s.chart {
		if (a.Y < coord.Y) || (a.Y == coord.Y && a.X < coord.X) {
			if s.clearPath(coord, a) {
				s.chart[coord] = s.chart[coord] + 1
				s.chart[a] = s.chart[a] + 1
			}
		}
	}
}

func (s *starmap) printBoard() {
	if debugPrint {
		fmt.Println()

		for i := 0; i < s.maxY; i++ {
			for j := 0; j < s.maxX; j++ {
				c := asCoord(j, i)
				a, ok := s.chart[c]
				if ok {
					fmt.Printf("%4d", a)
				} else {
					fmt.Print("    ")
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}
}
