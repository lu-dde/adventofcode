package u1p

import (
	"github.com/lu-dde/adventofcode/internal/solver"
)

//Problems is a list of problems for day11
func Problems() []solver.Problem {

	solvers := []solver.Problem{
		solver.Problem{
			Key:       "1p1",
			Day:       "1",
			Part:      "1",
			Solve:     Solve1,
			InputFile: "solutions/u1p/input.txt",
		},
		solver.Problem{
			Key:       "1p2",
			Day:       "1",
			Part:      "2",
			Solve:     Solve2,
			InputFile: "solutions/u1p/input.txt",
		},
	}

	return solvers

}
