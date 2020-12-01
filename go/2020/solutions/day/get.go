package day

import (
	"github.com/lu-dde/adventofcode/internal/solver"
	"github.com/lu-dde/adventofcode/solutions/day/p1"
	"github.com/lu-dde/adventofcode/solutions/day/p2"
)

const day string = "DAY"

//Problems is a list of problems for DAY
func Problems() []solver.Problem {

	solvers := []solver.Problem{
		{
			Key:       day + "p1",
			Day:       day,
			Part:      "1",
			Solve:     p1.Solve1,
			InputFile: "solutions/u1/input.txt",
		},
		{
			Key:       day + "p2",
			Day:       day,
			Part:      "2",
			Solve:     p2.Solve2,
			InputFile: "solutions/u1/input.txt",
		},
	}

	return solvers

}
