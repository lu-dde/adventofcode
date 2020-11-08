package u5p

import (
	"github.com/lu-dde/adventofcode/internal/solver"
	"github.com/lu-dde/adventofcode/solutions/u5p/u51"
	"github.com/lu-dde/adventofcode/solutions/u5p/u52"
)

const (
	inputPath = "solutions/u5p"
	day       = "5"
)

//Problems is a list of problems for day
func Problems() []solver.Problem {

	solvers := []solver.Problem{
		solver.Problem{
			Key:       day + "p1",
			Day:       day,
			Part:      "1",
			Solve:     u51.Solve,
			InputFile: inputPath + "/input.txt",
		},
		solver.Problem{
			Key:       day + "p2",
			Day:       day,
			Part:      "2",
			Solve:     u52.Solve,
			InputFile: inputPath + "/input.txt",
		},
	}

	return solvers

}
