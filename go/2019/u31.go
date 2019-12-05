package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type pos struct {
	x int
	y int
	d int
	w int
}

//U31 is main proxy for solve, takes a string channel
func U31(p chan string, s chan string) {

	var steps = math.MaxInt32
	var visited = map[int]pos{}

	var wireID = 3
	for wire := range p {
		wireID--
		var position = pos{x: 0, y: 0, d: 0, w: wireID}

		cmds := strings.Split(wire, ",")
		for _, cmd := range cmds {
			for _, c := range position.move(cmd) {
				v, ok := visited[c.hash()]

				if ok {
					if c.w < v.w {
						d := c.d + v.d
						if d < steps {
							steps = d
						}
					}
				} else {
					save := c
					visited[c.hash()] = save
				}

			}
		}
	}
	fmt.Println()

	s <- fmt.Sprintf("Solution: %d", steps)
}

func absI(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (origo *pos) hash() int {
	return origo.x*10000 + origo.y
}

func (origo *pos) distance() int {
	return absI(origo.x) + absI(origo.y)
}

func (origo *pos) move(cmd string) []pos {
	var dx = 0
	var dy = 0
	command := cmd[0]
	switch command {
	case 'U':
		dx = 1
	case 'D':
		dx = -1
	case 'R':
		dy = 1
	case 'L':
		dy = -1
	}
	amount, _ := strconv.Atoi(cmd[1:])

	var visited = []pos{}
	for index := 0; index < amount; index++ {
		visit := pos{x: dx + origo.x, y: dy + origo.y, d: 1 + origo.d, w: origo.w}
		origo.x = visit.x
		origo.y = visit.y
		origo.d = visit.d
		visited = append(visited, visit)
	}

	return visited
}
