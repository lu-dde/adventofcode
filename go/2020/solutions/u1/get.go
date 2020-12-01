package u1

import (
	"github.com/lu-dde/adventofcode/internal/solver"
	"github.com/lu-dde/adventofcode/solutions/u1/p1"
)

const day string = "1"

//Problems is a list of problems for 1
func Problems() []solver.Problem {

	solvers := []solver.Problem{
		{
			Key:       day + "p1",
			Day:       day,
			Part:      "1",
			Solve:     p1.Solve1,
			InputFile: "solutions/u1/input.txt",
		},
	}

	return solvers

}
