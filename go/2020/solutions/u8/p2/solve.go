package p2

import (
	"fmt"

	"github.com/lu-dde/adventofcode/internal/handheld"
	"github.com/lu-dde/adventofcode/internal/ops"
)

//Solve is main proxy for solve, takes a string channel
func Solve(p chan string, s chan string) {
	var t = 0

	origin := handheld.GetConsole(p)

	opsLen := len(origin.Ops)

	machines := []*handheld.Console{&origin}

	for i, o := range origin.Ops {
		if o.Action == ops.NOP && o.Amount != 0 {
			c := make([]ops.Cmd, opsLen)
			copy(c, origin.Ops)
			c[i] = ops.Cmd{
				Action: ops.JMP,
				Amount: o.Amount,
			}
			machines = append(machines, &handheld.Console{
				Ops: c,
				Fin: origin.Fin,
			})
		}
		if o.Action == ops.JMP {
			c := make([]ops.Cmd, opsLen)
			copy(c, origin.Ops)
			c[i] = ops.Cmd{
				Action: ops.NOP,
				Amount: o.Amount,
			}
			machines = append(machines, &handheld.Console{
				Ops: c,
				Fin: origin.Fin,
			})
		}
	}

outer:
	for {
		for _, c := range machines {
			more, err := c.Exec()
			if !more && err == nil {
				t = c.Acc
				break outer
			}
		}
	}

	s <- fmt.Sprintf("Solution: %d", t)
}
