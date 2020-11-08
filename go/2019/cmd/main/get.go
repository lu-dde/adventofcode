package main

import (
	"github.com/lu-dde/adventofcode/internal/solver"
	"github.com/lu-dde/adventofcode/solutions/u10p"
	"github.com/lu-dde/adventofcode/solutions/u11p"
	"github.com/lu-dde/adventofcode/solutions/u12p"
	"github.com/lu-dde/adventofcode/solutions/u13p"
	"github.com/lu-dde/adventofcode/solutions/u1p"
	"github.com/lu-dde/adventofcode/solutions/u2p"
	"github.com/lu-dde/adventofcode/solutions/u3p"
	"github.com/lu-dde/adventofcode/solutions/u4p"
	"github.com/lu-dde/adventofcode/solutions/u5p"
	"github.com/lu-dde/adventofcode/solutions/u6p"
	"github.com/lu-dde/adventofcode/solutions/u7p"
	"github.com/lu-dde/adventofcode/solutions/u8p"
	"github.com/lu-dde/adventofcode/solutions/u9p"
)

//GetProblem fetch a problem
func getProblem(name string) *solver.Problem {
	var solvers = []solver.Problem{}

	solvers = append(solvers, u1p.Problems()...)
	solvers = append(solvers, u2p.Problems()...)
	solvers = append(solvers, u3p.Problems()...)
	solvers = append(solvers, u4p.Problems()...)
	solvers = append(solvers, u5p.Problems()...)
	solvers = append(solvers, u6p.Problems()...)
	solvers = append(solvers, u7p.Problems()...)
	solvers = append(solvers, u8p.Problems()...)
	solvers = append(solvers, u9p.Problems()...)
	solvers = append(solvers, u10p.Problems()...)
	solvers = append(solvers, u11p.Problems()...)
	solvers = append(solvers, u12p.Problems()...)
	solvers = append(solvers, u13p.Problems()...)

	for _, p := range solvers {
		if name == p.Key {
			return &p
		}
	}

	return nil
}
