package p1

import (
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

func newOps(line string) *ops {
	cc := strings.Fields(line)
	i, _ := strconv.Atoi(cc[1])
	return &ops{
		cmd:    cc[0],
		amount: i,
		called: 0,
	}
}

type console struct {
	ops []*ops
	pos int
	acc int
}

func (c *console) exec() *ops {
	current := c.next()
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

	return c.next()
}

func (c *console) next() *ops {
	return c.ops[c.pos]
}

func getConsole(input <-chan string) console {
	ops := []*ops{}

	for line := range input {
		ops = append(ops, newOps(line))
	}

	return console{
		ops: ops,
		pos: 0,
		acc: 0,
	}
}

//Solve is main proxy for solve, takes a string channel
func Solve(p chan string, s chan string) {
	var t = 0

	c := getConsole(p)

	for {
		next := c.next()

		//	fmt.Println(c.pos, next, c.acc)

		if next.called > 0 {
			break
		}
		c.exec()
	}

	t = c.acc

	s <- fmt.Sprintf("Solution: %d", t)
}
