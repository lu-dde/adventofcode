package main

import (
	"github.com/lu-dde/adventofcode/internal/solver"
	"github.com/lu-dde/adventofcode/solutions/u1"
	"github.com/lu-dde/adventofcode/solutions/u10"
	"github.com/lu-dde/adventofcode/solutions/u11"
	"github.com/lu-dde/adventofcode/solutions/u12"
	"github.com/lu-dde/adventofcode/solutions/u13"
	"github.com/lu-dde/adventofcode/solutions/u14"
	"github.com/lu-dde/adventofcode/solutions/u15"
	"github.com/lu-dde/adventofcode/solutions/u16"
	"github.com/lu-dde/adventofcode/solutions/u17"
	"github.com/lu-dde/adventofcode/solutions/u18"
	"github.com/lu-dde/adventofcode/solutions/u2"
	"github.com/lu-dde/adventofcode/solutions/u3"
	"github.com/lu-dde/adventofcode/solutions/u4"
	"github.com/lu-dde/adventofcode/solutions/u5"
	"github.com/lu-dde/adventofcode/solutions/u6"
	"github.com/lu-dde/adventofcode/solutions/u7"
	"github.com/lu-dde/adventofcode/solutions/u8"
	"github.com/lu-dde/adventofcode/solutions/u9"
)

//GetProblem fetch a problem
func getProblem(name string) *solver.Problem {
	var solvers = []solver.Problem{}

	solvers = append(solvers, u1.Problems()...)
	solvers = append(solvers, u2.Problems()...)
	solvers = append(solvers, u3.Problems()...)
	solvers = append(solvers, u4.Problems()...)
	solvers = append(solvers, u5.Problems()...)
	solvers = append(solvers, u6.Problems()...)
	solvers = append(solvers, u7.Problems()...)
	solvers = append(solvers, u8.Problems()...)
	solvers = append(solvers, u9.Problems()...)
	solvers = append(solvers, u10.Problems()...)
	solvers = append(solvers, u11.Problems()...)
	solvers = append(solvers, u12.Problems()...)
	solvers = append(solvers, u13.Problems()...)
	solvers = append(solvers, u14.Problems()...)
	solvers = append(solvers, u15.Problems()...)
	solvers = append(solvers, u16.Problems()...)
	solvers = append(solvers, u17.Problems()...)
	solvers = append(solvers, u18.Problems()...)

	for _, p := range solvers {
		if name == p.Key {
			return &p
		}
	}

	return nil
}
