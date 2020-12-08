package handheld

import (
	"errors"

	"github.com/lu-dde/adventofcode/internal/ops"
)

//Console struct
type Console struct {
	Ops    []ops.Cmd
	Pos    int
	Acc    int
	Fin    int
	Exited bool
}

//Exec the next Ops or exit
func (c *Console) Exec() (bool, error) {
	more, err := c.HasMore()

	if !more {
		return more, err
	}

	current := &c.Ops[c.Pos]
	current.Called++

	switch current.Action {
	case ops.NOP:
		c.Pos++
	case ops.ACC:
		c.Acc += current.Amount
		c.Pos++
	case ops.JMP:
		c.Pos += current.Amount
	}
	return c.HasMore()
}

//HasMore Ops or not
func (c *Console) HasMore() (bool, error) {
	if c.Pos == c.Fin {
		c.Exited = true
		return false, nil
	}
	if c.Pos > c.Fin {
		c.Exited = true
		return false, errors.New("out of bounds")
	}
	return true, nil
}
