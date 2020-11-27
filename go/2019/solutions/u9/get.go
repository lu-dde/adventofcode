package u9

import (
	"github.com/lu-dde/adventofcode/internal/solver"
	"github.com/lu-dde/adventofcode/solutions/u9/u91"
	"github.com/lu-dde/adventofcode/solutions/u9/u92"
)

const (
	inputPath = "solutions/u9"
	day       = "9"
)

//Problems is a list of problems for day
func Problems() []solver.Problem {

	solvers := []solver.Problem{
		{
			Key:       day + "p1",
			Day:       day,
			Part:      "1",
			Solve:     u91.Solve,
			InputFile: inputPath + "/input.txt",
		},
		{
			Key:       day + "p2",
			Day:       day,
			Part:      "2",
			Solve:     u92.Solve,
			InputFile: inputPath + "/input.txt",
		},
	}

	return solvers

}
