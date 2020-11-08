package u4p

import (
	"github.com/lu-dde/adventofcode/internal/solver"
	"github.com/lu-dde/adventofcode/solutions/u4p/u41"
	"github.com/lu-dde/adventofcode/solutions/u4p/u42"
)

const (
	inputPath = "solutions/u4p"
	day       = "4"
)

//Problems is a list of problems for day
func Problems() []solver.Problem {

	solvers := []solver.Problem{
		solver.Problem{
			Key:       day + "p1",
			Day:       day,
			Part:      "1",
			Solve:     u41.Solve,
			InputFile: inputPath + "/input.txt",
		},
		solver.Problem{
			Key:       day + "p1t1",
			Day:       day,
			Part:      "1",
			Solve:     u41.Solve,
			InputFile: inputPath + "/test.1.txt",
		},
		solver.Problem{
			Key:       day + "p1t1",
			Day:       day,
			Part:      "1",
			Solve:     u41.Solve,
			InputFile: inputPath + "/test.1955.txt",
		},
		solver.Problem{
			Key:       day + "p2",
			Day:       day,
			Part:      "2",
			Solve:     u42.Solve,
			InputFile: inputPath + "/input.txt",
		},
	}

	return solvers

}
