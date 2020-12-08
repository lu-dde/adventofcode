package p2

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	nop = "nop"
	jmp = "jmp"
	acc = "acc"
)

type ops struct {
	cmd    string
	amount int
	called int
}

func newOps(line string) ops {
	cc := strings.Fields(line)
	i, _ := strconv.Atoi(cc[1])
	return ops{
		cmd:    cc[0],
		amount: i,
	}
}

type console struct {
	ops  []ops
	pos  int
	acc  int
	fin  int
	exit bool
}

func (c *console) exec() (bool, error) {
	more, err := c.hasNext()

	if !more {
		return more, err
	}

	current := &c.ops[c.pos]
	current.called++

	switch current.cmd {
	case nop:
		c.pos++
	case acc:
		c.acc += current.amount
		c.pos++
	case jmp:
		c.pos += current.amount
	}
	return c.hasNext()
}

func (c *console) hasNext() (bool, error) {
	if c.pos == c.fin {
		c.exit = true
		return false, nil
	}
	if c.pos > c.fin {
		c.exit = true
		return false, errors.New("out of bounds")
	}
	return true, nil
}

func getConsole(input <-chan string) console {
	ops := []ops{}

	for line := range input {
		ops = append(ops, newOps(line))
	}

	return console{
		ops: ops,
		fin: len(ops),
	}
}

//Solve is main proxy for solve, takes a string channel
func Solve(p chan string, s chan string) {
	var t = 0

	origin := getConsole(p)

	opsLen := len(origin.ops)

	machines := []*console{&origin}

	for i, o := range origin.ops {
		if o.cmd == nop && o.amount != 0 {
			c := make([]ops, opsLen)
			copy(c, origin.ops)
			c[i] = ops{
				cmd:    jmp,
				amount: o.amount,
			}
			machines = append(machines, &console{
				ops: c,
				fin: origin.fin,
			})
		}
		if o.cmd == jmp {
			c := make([]ops, opsLen)
			copy(c, origin.ops)
			c[i] = ops{
				cmd:    nop,
				amount: o.amount,
			}
			machines = append(machines, &console{
				ops: c,
				fin: origin.fin,
			})
		}
	}

outer:
	for {
		//fmt.Println("step")
		for _, c := range machines {
			more, err := c.exec()
			if !more && err == nil {
				t = c.acc
				break outer
			}
		}
		//fmt.Println()

	}

	/*
		for _, c := range machines {
			fmt.Println(c.exit, c.pos, c.fin, c.acc)
		}
	*/

	s <- fmt.Sprintf("Solution: %d", t)
}
