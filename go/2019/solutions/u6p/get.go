package u6p

import (
	"github.com/lu-dde/adventofcode/internal/solver"
	"github.com/lu-dde/adventofcode/solutions/u6p/u61"
	"github.com/lu-dde/adventofcode/solutions/u6p/u62"
)

const (
	inputPath = "solutions/u6p"
	day       = "6"
)

//Problems is a list of problems for day
func Problems() []solver.Problem {

	solvers := []solver.Problem{
		solver.Problem{
			Key:       day + "p1",
			Day:       day,
			Part:      "1",
			Solve:     u61.Solve,
			InputFile: inputPath + "/input.txt",
		},
		solver.Problem{
			Key:       day + "p2",
			Day:       day,
			Part:      "2",
			Solve:     u62.Solve,
			InputFile: inputPath + "/input.txt",
		},
	}

	return solvers

}
