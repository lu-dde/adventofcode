package u2

import (
	"github.com/lu-dde/adventofcode/internal/solver"
	"github.com/lu-dde/adventofcode/solutions/u2/p1"
	"github.com/lu-dde/adventofcode/solutions/u2/p2"
)

const day string = "2"

//Problems is a list of problems for DAY
func Problems() []solver.Problem {

	path := "solutions/u" + day

	solvers := []solver.Problem{
		{
			Key:       day + "p1",
			Day:       day,
			Part:      "1",
			Solve:     p1.Solve1,
			InputFile: path + "/input.txt",
		},
		{
			Key:       day + "p1t1",
			Day:       day,
			Part:      "1",
			Solve:     p1.Solve1,
			InputFile: path + "/test.2.txt",
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
			InputFile: path + "/test.2.txt",
		},
	}

	return solvers

}
