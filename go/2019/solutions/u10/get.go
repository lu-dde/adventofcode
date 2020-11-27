package u10

import (
	"github.com/lu-dde/adventofcode/internal/solver"
	"github.com/lu-dde/adventofcode/solutions/u10/u101"
	"github.com/lu-dde/adventofcode/solutions/u10/u102"
)

const (
	inputPath = "solutions/u10"
	day       = "10"
)

//Problems is a list of problems for day
func Problems() []solver.Problem {

	solvers := []solver.Problem{
		{
			Key:       day + "p1",
			Day:       day,
			Part:      "1",
			Solve:     u101.Solve,
			InputFile: inputPath + "/input.txt",
		},
		{
			Key:       day + "p2",
			Day:       day,
			Part:      "2",
			Solve:     u102.Solve,
			InputFile: inputPath + "/input.txt",
		},
	}

	return solvers

}
