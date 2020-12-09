package u9

import (
	"github.com/lu-dde/adventofcode/internal/solver"
	"github.com/lu-dde/adventofcode/solutions/u9/p1"
	"github.com/lu-dde/adventofcode/solutions/u9/p2"
)

const day string = "9"

//Problems is a list of problems for 9
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
			Key:       day + "p1t1",
			Day:       day,
			Part:      "1",
			Solve:     p1.Solve,
			InputFile: path + "/short.txt",
		},
		{
			Key:       day + "p2",
			Day:       day,
			Part:      "2",
			Solve:     p2.Solve,
			InputFile: path + "/input.txt",
		},
	}

	return solvers

}
