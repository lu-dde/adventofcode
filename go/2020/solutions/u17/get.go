package u17

import (
	"github.com/lu-dde/adventofcode/internal/solver"
	"github.com/lu-dde/adventofcode/solutions/u17/p1"
	"github.com/lu-dde/adventofcode/solutions/u17/p2"
)

const day string = "17"

//Problems is a list of problems for 17
func Problems() []solver.Problem {

	path := "solutions/u" + day

	solvers := []solver.Problem{
		{
			Key:       day + "p1",
			Day:       day,
			Part:      "1",
			Solve:     p1.Solve,
			InputFile: path + "/input.txt",
		},
		{
			Key:       day + "p1t0",
			Day:       day,
			Part:      "1",
			Solve:     p1.Solve,
			InputFile: path + "/test.0.txt",
		},
		{
			Key:       day + "p1t1",
			Day:       day,
			Part:      "1",
			Solve:     p1.Solve,
			InputFile: path + "/test.112.txt",
		},
		{
			Key:       day + "p2",
			Day:       day,
			Part:      "2",
			Solve:     p2.Solve,
			InputFile: path + "/input.txt",
		},
		{
			Key:       day + "p2t1",
			Day:       day,
			Part:      "2",
			Solve:     p2.Solve,
			InputFile: path + "/test.112.txt",
		},
	}

	return solvers

}
