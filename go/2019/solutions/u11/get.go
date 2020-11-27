package u11

import (
	"github.com/lu-dde/adventofcode/internal/solver"
	p1 "github.com/lu-dde/adventofcode/solutions/u11/p1"
	p2 "github.com/lu-dde/adventofcode/solutions/u11/p2"
)

//Problems is a list of problems for day11
func Problems() []solver.Problem {

	solvers := []solver.Problem{
		{
			Key:       "11p1",
			Day:       "11",
			Part:      "1",
			Solve:     p1.Solve,
			InputFile: "solutions/u11/input/input.txt",
		},
		{
			Key:       "11p2",
			Day:       "11",
			Part:      "2",
			Solve:     p2.Solve,
			InputFile: "solutions/u11/input/input.txt",
		},
	}

	return solvers

}
