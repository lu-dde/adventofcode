package u10p

import (
	"github.com/lu-dde/adventofcode/internal/solver"
	"github.com/lu-dde/adventofcode/solutions/u10p/u101"
	"github.com/lu-dde/adventofcode/solutions/u10p/u102"
)

const (
	inputPath = "solutions/u10p"
	day       = "10"
)

//Problems is a list of problems for day
func Problems() []solver.Problem {

	solvers := []solver.Problem{
		solver.Problem{
			Key:       day + "p1",
			Day:       day,
			Part:      "1",
			Solve:     u101.Solve,
			InputFile: inputPath + "/input.txt",
		},
		solver.Problem{
			Key:       day + "p2",
			Day:       day,
			Part:      "2",
			Solve:     u102.Solve,
			InputFile: inputPath + "/input.txt",
		},
	}

	return solvers

}
