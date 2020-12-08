package ops

import (
	"strconv"
	"strings"
)

const (
	//NOP cmd
	NOP = "nop"
	//JMP cmd
	JMP = "jmp"
	//ACC cmd
	ACC = "acc"
)

//Cmd with amount, called
type Cmd struct {
	Action string
	Amount int
	Called int
}

//New from a string
func New(line string) Cmd {
	cc := strings.Fields(line)
	i, _ := strconv.Atoi(cc[1])
	return Cmd{
		Action: cc[0],
		Amount: i,
	}
}
