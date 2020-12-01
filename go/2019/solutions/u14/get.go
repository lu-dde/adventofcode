package u14

import (
	"github.com/lu-dde/adventofcode/internal/solver"
	"github.com/lu-dde/adventofcode/solutions/u14/p1"
)

const (
	inputPath = "solutions/u14"
	day       = "14"
)

//Problems is a list of problems for day
func Problems() []solver.Problem {

	solvers := []solver.Problem{
		{
			Key:       day + "p1",
			Day:       day,
			Part:      "1",
			Solve:     p1.Solve,
			InputFile: inputPath + "/input.txt",
		},
		{
			Key:       day + "p1t0",
			Day:       day,
			Part:      "1",
			Solve:     p1.Solve,
			InputFile: inputPath + "/test.1.165.txt",
		},
		{
			Key:       day + "p1t1",
			Day:       day,
			Part:      "1",
			Solve:     p1.Solve,
			InputFile: inputPath + "/test.1.13312.txt",
		},
		{
			Key:       day + "p1t2",
			Day:       day,
			Part:      "1",
			Solve:     p1.Solve,
			InputFile: inputPath + "/test.1.180697.txt",
		},
		{
			Key:       day + "p1t3",
			Day:       day,
			Part:      "1",
			Solve:     p1.Solve,
			InputFile: inputPath + "/test.1.2210736.txt",
		},
	}

	return solvers

}
