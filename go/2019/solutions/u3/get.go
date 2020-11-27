package u3

import (
	"github.com/lu-dde/adventofcode/internal/solver"
)

//Problems is a list of problems for day11
func Problems() []solver.Problem {

	solvers := []solver.Problem{
		{
			Key:       "3p1",
			Day:       "3",
			Part:      "1",
			Solve:     Solve,
			InputFile: "solutions/u3/input.txt",
		},
		{
			Key:       "3p1t1",
			Day:       "3",
			Part:      "1",
			Solve:     Solve,
			InputFile: "solutions/u3/test.1.txt",
		},
	}

	return solvers

}
