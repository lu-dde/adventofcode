package u4

import (
	"github.com/lu-dde/adventofcode/internal/solver"
	"github.com/lu-dde/adventofcode/solutions/u4/u41"
	"github.com/lu-dde/adventofcode/solutions/u4/u42"
)

const (
	inputPath = "solutions/u4"
	day       = "4"
)

//Problems is a list of problems for day
func Problems() []solver.Problem {

	solvers := []solver.Problem{
		{
			Key:       day + "p1",
			Day:       day,
			Part:      "1",
			Solve:     u41.Solve,
			InputFile: inputPath + "/input.txt",
		},
		{
			Key:       day + "p1t1",
			Day:       day,
			Part:      "1",
			Solve:     u41.Solve,
			InputFile: inputPath + "/test.1.txt",
		},
		{
			Key:       day + "p1t1",
			Day:       day,
			Part:      "1",
			Solve:     u41.Solve,
			InputFile: inputPath + "/test.1955.txt",
		},
		{
			Key:       day + "p2",
			Day:       day,
			Part:      "2",
			Solve:     u42.Solve,
			InputFile: inputPath + "/input.txt",
		},
	}

	return solvers

}
