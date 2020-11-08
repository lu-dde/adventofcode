package u3p

import (
	"github.com/lu-dde/adventofcode/internal/solver"
)

//Problems is a list of problems for day11
func Problems() []solver.Problem {

	solvers := []solver.Problem{
		solver.Problem{
			Key:       "3p1",
			Day:       "3",
			Part:      "1",
			Solve:     Solve,
			InputFile: "solutions/u3p/input.txt",
		},
		solver.Problem{
			Key:       "3p1t1",
			Day:       "3",
			Part:      "1",
			Solve:     Solve,
			InputFile: "solutions/u3p/test.1.txt",
		},
	}

	return solvers

}
