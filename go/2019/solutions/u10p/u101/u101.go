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

	stars := Starmap{
		Chart: make(map[coord.Pair]int, size*size),
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
				stars.AddAstroid(asCoord(x, y))
			}
		}
		y++
	}

	stars.SetRayMap()

	if debugPrint {
		for row := 0; row < stars.MaxY; row++ {
			for column := 0; column < stars.MaxX; column++ {
				ray := asCoord(row, column)
				distance, ok := stars.Rays.Get(ray)
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

		stars.PrintBoard()
	}

	maxCoord, maxValue := stars.GetU101()

	s <- fmt.Sprintf("Best is %d,%d with %d other asteroids detected.", maxCoord.X, maxCoord.Y, maxValue)
}

func asCoord(x, y int) coord.Pair {
	return coord.NewPair(x, y)
}

type Starmap struct {
	Chart map[coord.Pair]int
	MaxX  int
	MaxY  int
	Rays  rays.Map
}

func (s *Starmap) AddAstroid(coord coord.Pair) {
	s.Chart[coord] = 0
	s.extend(coord)
}

func (s *Starmap) extend(coord coord.Pair) {
	if s.MaxX < coord.X {
		s.MaxX = coord.X
	}
	if s.MaxY < coord.Y {
		s.MaxY = coord.Y
	}
}

func (s *Starmap) SetRayMap() {
	small, large := s.MaxX, s.MaxY
	if small > large {
		small, large = large, small
	}
	s.Rays = rays.New(large)
}

func (s *Starmap) look() {
	for a := range s.Chart {
		s.lookNorth(a)
	}
}

func (s *Starmap) clearPath(goal, other coord.Pair) bool {
	distance, direction := goal.Distance(other)
	ray, _ := s.Rays.Get(distance)
	speed := ray.Speed.ChangeDirection(direction)

	for {
		other = other.Add(speed)
		if goal == other {
			break
		}
		if _, ok := s.Chart[other]; ok {
			return false
		}
	}

	return true
}

func (s *Starmap) lookNorth(coord coord.Pair) {
	for a := range s.Chart {
		if (a.Y < coord.Y) || (a.Y == coord.Y && a.X < coord.X) {
			if s.clearPath(coord, a) {
				s.Chart[coord] = s.Chart[coord] + 1
				s.Chart[a] = s.Chart[a] + 1
			}
		}
	}
}

func (s *Starmap) GetU101() (coord.Pair, int) {

	s.look()

	var maxCoord coord.Pair
	maxValue := 0
	for coord, value := range s.Chart {
		if value > maxValue {
			maxValue = value
			maxCoord = coord
		}
	}

	return maxCoord, maxValue
}
func (s *Starmap) PrintBoard() {
	fmt.Println("######")

	for i := 0; i <= s.MaxY; i++ {
		for j := 0; j <= s.MaxX; j++ {
			c := asCoord(j, i)
			a, ok := s.Chart[c]
			if ok {
				fmt.Printf("%3d", a)
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Println()
	}
	fmt.Println("######")
}
