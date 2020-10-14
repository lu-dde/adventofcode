package main

import (
	"github.com/lu-dde/adventofcode/solutions/u101"
	"github.com/lu-dde/adventofcode/solutions/u11"
	"github.com/lu-dde/adventofcode/solutions/u12"
	"github.com/lu-dde/adventofcode/solutions/u21"
	"github.com/lu-dde/adventofcode/solutions/u22"
	"github.com/lu-dde/adventofcode/solutions/u31"
	"github.com/lu-dde/adventofcode/solutions/u41"
	"github.com/lu-dde/adventofcode/solutions/u42"
	"github.com/lu-dde/adventofcode/solutions/u51"
	"github.com/lu-dde/adventofcode/solutions/u52"
	"github.com/lu-dde/adventofcode/solutions/u61"
	"github.com/lu-dde/adventofcode/solutions/u62"
	"github.com/lu-dde/adventofcode/solutions/u71"
	"github.com/lu-dde/adventofcode/solutions/u72"
	"github.com/lu-dde/adventofcode/solutions/u81"
	"github.com/lu-dde/adventofcode/solutions/u82"
	"github.com/lu-dde/adventofcode/solutions/u91"
	"github.com/lu-dde/adventofcode/solutions/u92"
)

type solver func(chan string, chan string)

func getSolver(name string) solver {
	solver, ok := map[string]solver{
		"11":  u11.Solve,
		"12":  u12.Solve,
		"21":  u21.Solve,
		"22":  u22.Solve,
		"31":  u31.Solve,
		"41":  u41.Solve,
		"42":  u42.Solve,
		"51":  u51.Solve,
		"52":  u52.Solve,
		"61":  u61.Solve,
		"62":  u62.Solve,
		"71":  u71.Solve,
		"72":  u72.Solve,
		"81":  u81.Solve,
		"82":  u82.Solve,
		"91":  u91.Solve,
		"92":  u92.Solve,
		"101": u101.Solve,
	}[name]

	if !ok {
		panic("could not find a solver function")
	}

	return solver
}

func getTestfile(name string) string {
	textfile, ok := map[string]string{
		"11":  "input/u11.txt",
		"12":  "input/u11.txt",
		"21":  "input/u21.txt",
		"22":  "input/u21.txt",
		"31":  "input/u31.txt",
		"41":  "input/u41.txt",
		"51":  "input/u51.txt",
		"52":  "input/u51.txt",
		"61":  "input/u61.txt",
		"62":  "input/u61.txt",
		"71":  "input/u71.txt",
		"72":  "input/u71.txt",
		"81":  "input/u81.txt",
		"82":  "input/u81.txt",
		"91":  "input/u91.txt",
		"92":  "input/u91.txt",
		"101": "input/u101.test.1.2.35.txt",
	}[name]

	if !ok {
		panic("could not find a solver function")
	}

	return textfile
}
