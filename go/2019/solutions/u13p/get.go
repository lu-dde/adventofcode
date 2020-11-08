package u13p

import (
	"github.com/lu-dde/adventofcode/internal/solver"
	"github.com/lu-dde/adventofcode/solutions/u13p/p1"
	"github.com/lu-dde/adventofcode/solutions/u13p/p2"
)

const (
	inputPath = "solutions/u13p"
	day       = "13"
)

//Problems is a list of problems for day
func Problems() []solver.Problem {

	solvers := []solver.Problem{
		solver.Problem{
			Key:       day + "p1",
			Day:       day,
			Part:      "1",
			Solve:     p1.Solve,
			InputFile: inputPath + "/input.txt",
		},
		solver.Problem{
			Key:       day + "p2",
			Day:       day,
			Part:      "2",
			Solve:     p2.Solve,
			InputFile: inputPath + "/input.txt",
		},
	}

	return solvers

}
