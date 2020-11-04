package u102

import (
	"fmt"
	"sort"

	"github.com/lu-dde/adventofcode/internal/coord"
	"github.com/lu-dde/adventofcode/solutions/u101"
)

//Solve U102
func Solve(p chan string, s chan string) {

	size := 40

	stars := u101.Starmap{
		Chart: make(map[coord.Pair]int, size*size),
	}

	/*
		.#....#####...#..
		##...##.#####..##
		##...#...#.#####.
		..#.....#...###..
		...#.....#....##
	*/
	row := 0
	for line := range p {
		for column, t := range line {
			if t == '#' {
				stars.AddAstroid(coord.NewPair(column, row))
			}
		}
		//fmt.Println(row, line)
		row++
	}

	laser := LaserMap{
		Looks: stars,
	}

	laser.Init()

	fmt.Printf("Best Station at (%d,%d).\n", laser.Station.X, laser.Station.Y)
	a := laser.Targets[199]

	s <- fmt.Sprintf("The 200th asteroid to be vaporized is at (%d,%d) with score %d", a.target.X, a.target.Y, a.target.X*100+a.target.Y)
}

type LaserMap struct {
	Looks   u101.Starmap
	Station coord.Pair
	Targets laserSlice
}

func (l *LaserMap) Init() {
	l.Looks.SetRayMap()
	l.Station, _ = l.Looks.GetU101()

	//l.Looks.PrintBoard()

	l.Fire()

	sort.Stable(l.Targets)
}

func (l *LaserMap) Fire() {
	for target := range l.Looks.Chart {
		l.LaserDistance(target)
	}
}

func (l *LaserMap) LaserDistance(target coord.Pair) {
	origin := target
	if origin == l.Station {
		return
	}

	distance, direction := l.Station.Distance(target)
	ray, _ := l.Looks.Rays.Get(distance)
	speed := ray.Speed.ChangeDirection(direction)
	returnSpeed := speed.ChangeDirection(coord.Pair{X: -1, Y: -1})
	crashes := 0
	for l.Station != target {
		target = target.Add(speed)
		if _, ok := l.Looks.Chart[target]; ok {
			crashes++
		}
	}
	l.Targets = append(l.Targets, beam{
		target:  origin,
		crashes: crashes,
		angle:   angle(returnSpeed),
	})
	//fmt.Println("O", origin, "dir", speed, "\tres", "crashes", crashes, "angle", angle(returnSpeed)/(2*math.Pi)*360)
}

type beam struct {
	target  coord.Pair
	crashes int
	angle   float64
}

type laserSlice []beam

func (cs laserSlice) Len() int {
	return len(cs)
}
func (cs laserSlice) Less(i, j int) bool {
	p, q := cs[i], cs[j]
	lowerGeneration := p.crashes < q.crashes
	sameGeneration := p.crashes == q.crashes
	lowerAngle := p.angle < q.angle

	return lowerGeneration || (sameGeneration && lowerAngle)
}
func (cs laserSlice) Swap(i, j int) {
	cs[i], cs[j] = cs[j], cs[i]
}
