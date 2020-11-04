package u11p

import (
	"github.com/lu-dde/adventofcode/internal/solver"
	"github.com/lu-dde/adventofcode/solutions/u11p/u11p1"
	"github.com/lu-dde/adventofcode/solutions/u11p/u11p2"
)

//Problems is a list of problems for day11
func Problems() []solver.Problem {

	solvers := []solver.Problem{
		solver.Problem{
			Key:       "11p1",
			Day:       "11",
			Part:      "1",
			Solve:     u11p1.Solve,
			InputFile: "solutions/u11p/input/input.txt",
		},
		solver.Problem{
			Key:       "11p2",
			Day:       "11",
			Part:      "2",
			Solve:     u11p2.Solve,
			InputFile: "solutions/u11p/input/input.txt",
		},
	}

	return solvers

}
