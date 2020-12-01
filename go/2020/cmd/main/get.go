package main

import (
	"github.com/lu-dde/adventofcode/internal/solver"
	"github.com/lu-dde/adventofcode/solutions/u1"
)

//GetProblem fetch a problem
func getProblem(name string) *solver.Problem {
	var solvers = []solver.Problem{}

	solvers = append(solvers, u1.Problems()...)

	for _, p := range solvers {
		if name == p.Key {
			return &p
		}
	}

	return nil
}
