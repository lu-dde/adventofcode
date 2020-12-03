package main

import (
	"github.com/lu-dde/adventofcode/internal/solver"
	"github.com/lu-dde/adventofcode/solutions/u1"
	"github.com/lu-dde/adventofcode/solutions/u2"
	"github.com/lu-dde/adventofcode/solutions/u3"
)

//GetProblem fetch a problem
func getProblem(name string) *solver.Problem {
	var solvers = []solver.Problem{}

	solvers = append(solvers, u1.Problems()...)
	solvers = append(solvers, u2.Problems()...)
	solvers = append(solvers, u3.Problems()...)

	for _, p := range solvers {
		if name == p.Key {
			return &p
		}
	}

	return nil
}
