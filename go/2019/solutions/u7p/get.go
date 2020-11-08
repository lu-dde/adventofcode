package u7p

import (
	"github.com/lu-dde/adventofcode/internal/solver"
	"github.com/lu-dde/adventofcode/solutions/u7p/u71"
	"github.com/lu-dde/adventofcode/solutions/u7p/u72"
)

const (
	inputPath = "solutions/u7p"
	day       = "7"
)

//Problems is a list of problems for day
func Problems() []solver.Problem {

	solvers := []solver.Problem{
		solver.Problem{
			Key:       day + "p1",
			Day:       day,
			Part:      "1",
			Solve:     u71.Solve,
			InputFile: inputPath + "/input.txt",
		},
		solver.Problem{
			Key:       day + "p2",
			Day:       day,
			Part:      "2",
			Solve:     u72.Solve,
			InputFile: inputPath + "/input.txt",
		},
	}

	return solvers

}
