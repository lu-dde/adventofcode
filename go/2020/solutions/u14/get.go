package u14

import (
	"github.com/lu-dde/adventofcode/internal/solver"
	"github.com/lu-dde/adventofcode/solutions/u14/p1"
	"github.com/lu-dde/adventofcode/solutions/u14/p2"
)

const day string = "14"

//Problems is a list of problems for 14
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
			InputFile: path + "/test.165.txt",
		},
		{
			Key:       day + "p2",
			Day:       day,
			Part:      "2",
			Solve:     p2.Solve,
			InputFile: path + "/input.txt",
		},
		{
			Key:       day + "p2t0",
			Day:       day,
			Part:      "2",
			Solve:     p2.Solve,
			InputFile: path + "/test.208.txt",
		},
	}

	return solvers

}
