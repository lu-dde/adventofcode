package u8p

import (
	"github.com/lu-dde/adventofcode/internal/solver"
	"github.com/lu-dde/adventofcode/solutions/u8p/u81"
	"github.com/lu-dde/adventofcode/solutions/u8p/u82"
)

const (
	inputPath = "solutions/u8p"
	day       = "8"
)

//Problems is a list of problems for day
func Problems() []solver.Problem {

	solvers := []solver.Problem{
		solver.Problem{
			Key:       day + "p1",
			Day:       day,
			Part:      "1",
			Solve:     u81.Solve,
			InputFile: inputPath + "/input.txt",
		},
		solver.Problem{
			Key:       day + "p2",
			Day:       day,
			Part:      "2",
			Solve:     u82.Solve,
			InputFile: inputPath + "/input.txt",
		},
	}

	return solvers

}
