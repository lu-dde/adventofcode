package u2

import (
	"github.com/lu-dde/adventofcode/internal/solver"
	"github.com/lu-dde/adventofcode/solutions/u2/u21"
	"github.com/lu-dde/adventofcode/solutions/u2/u22"
)

//Problems is a list of problems for day11
func Problems() []solver.Problem {

	solvers := []solver.Problem{
		{
			Key:       "2p1",
			Day:       "2",
			Part:      "1",
			Solve:     u21.Solve,
			InputFile: "solutions/u2/input.txt",
		},
		{
			Key:       "2p2",
			Day:       "2",
			Part:      "2",
			Solve:     u22.Solve,
			InputFile: "solutions/u2/input.txt",
		},
	}

	return solvers

}
