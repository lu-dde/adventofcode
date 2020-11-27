package u12

import (
	"github.com/lu-dde/adventofcode/internal/solver"
	"github.com/lu-dde/adventofcode/solutions/u12/p1"
	"github.com/lu-dde/adventofcode/solutions/u12/p2"
)

//Problems is a list of problems for day11
func Problems() []solver.Problem {

	solvers := []solver.Problem{
		{
			Key:       "12p1",
			Day:       "12",
			Part:      "1",
			Solve:     p1.Solve,
			InputFile: "solutions/u12/input/input.txt",
		},
		{
			Key:       "12p1t1",
			Day:       "12",
			Part:      "1",
			Solve:     p1.Solve,
			InputFile: "solutions/u12/input/test.1.txt",
		},
		{
			Key:       "12p1t2",
			Day:       "12",
			Part:      "1",
			Solve:     p1.Solve,
			InputFile: "solutions/u12/input/test.2.txt",
		},
		{
			Key:       "12p2",
			Day:       "12",
			Part:      "2",
			Solve:     p2.Solve,
			InputFile: "solutions/u12/input/input.txt",
		},
		{
			Key:       "12p2t1",
			Day:       "12",
			Part:      "2",
			Solve:     p2.Solve,
			InputFile: "solutions/u12/input/test.1.txt",
		},
		{
			Key:       "12p2t2",
			Day:       "12",
			Part:      "2",
			Solve:     p2.Solve,
			InputFile: "solutions/u12/input/test.2.txt",
		},
	}

	return solvers

}
